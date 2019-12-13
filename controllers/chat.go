package controllers

import (
	"github.com/gorilla/websocket"
	"time"
	"sync"
	"net/http"
	"CustomIM/services"
	"strconv"
	"github.com/astaxie/beego/logs"
	"CustomIM/utils"
	"CustomIM/models"
	"github.com/go-errors/errors"
	"io"
	"encoding/json"
)

//每一个用户对应一个连接节点
type Node struct {
	Conn 	*websocket.Conn
	Data 	chan []byte
	Heart 	*time.Timer
	Close 	chan bool
	GroupId 	int
	SubId 		int
}
//关系映射表
var GroupClientMap  = make(map[int]map[int]*Node, 0)
//读写锁
var rwlocker sync.RWMutex
//todo websocket通讯接口
func Chat(w http.ResponseWriter, r *http.Request) {
	//todo 获取请求的参数
	params := r.URL.Query()
	role := params.Get("role")
	uuid := params.Get("uuid")
	id, _ := strconv.Atoi(params.Get("id"))
	var (
		ip 			string
		groupid 	int
		err 		error
		ipuser		*models.IpUsers
		app 		*models.Apps
	)
	//todo 检验接入是否合法
	if role == "admin" {
		token := params.Get("token")
		groupid, err = services.CheckChat(id, token, uuid)
		if err != nil {
			logs.Error(err.Error())
			return
		}
	} else if role == "ip" {
		app,err = services.GetAppByIdAndUuid(id,uuid)
		if err != nil {
			logs.Error(err.Error())
			return
		}
		//获取请求的ip地址
		ip  = utils.ParseAddr(r.RemoteAddr)
		//判断是否存在Ip用户，没有则创建
		ipuser, err = services.FindOrCreateIpUser(app.Id,app.Uid, ip)
		if err != nil {
			logs.Error(err.Error())
			return
		}
		//更新链接时间
		err := services.UpdateConnect(ipuser.Id)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	//todo 创建连接节点
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	node := &Node{
		Conn:	conn,
		Data: 	make(chan []byte, 50),
		Heart:	time.NewTimer(time.Minute),
		Close: 	make(chan bool),
	}
	//todo 将键名与节点进行绑定
	if role == "admin" {
		node.GroupId = groupid
		node.SubId = 0
		//写入关系，加索
		if GroupClientMap[groupid] == nil {
			GroupClientMap[groupid] = make(map[int]*Node)
		}
		rwlocker.Lock()
		GroupClientMap[groupid][0] = node
		rwlocker.Unlock()
	} else if role == "ip" {
		node.GroupId = ipuser.Uid
		node.SubId = ipuser.Id
		if GroupClientMap[ipuser.Uid] == nil {
			GroupClientMap[ipuser.Uid] = make(map[int]*Node)
		}
		rwlocker.Lock()
		GroupClientMap[ipuser.Uid][ipuser.Id] = node
		rwlocker.Unlock()
	}
	//todo 开启监听websocket协程
	go ListenWebSocket(node)
	//todo 开启监听NODE协程
	go ListenNode(node)
}
//todo 监听websocket消息
func ListenWebSocket(node *Node) {
	for {
		//监听消息
		_, msg, err := node.Conn.ReadMessage()
		if err != nil {
			//断开连接
			logs.Error("断开连接",err.Error())
			//结束监听Node的协程
			node.Close <- true
			//删除关系表的节点
			if err := deleteNode(node.GroupId, node.SubId); err != nil {
				logs.Error(err.Error())
				return
			}

			return
		}
		//处理消息
		parseMsg(node,[]byte(msg))
	}
}
//todo 监听Node的消息
func ListenNode(node *Node) {
	for {
		select {
			case data := <- node.Data:
				//发送消息
				err := node.Conn.WriteMessage(websocket.TextMessage,data)
				if err != nil {
					logs.Error(err.Error())
					return
				}
				//发送消息，重置计时器
				node.Heart.Reset(time.Minute)
			case <- node.Close:
				return
			case <- node.Heart.C:
				//心跳机制
				node.Conn.WriteMessage(websocket.TextMessage, []byte("heart"))
				node.Heart.Reset(time.Minute)

		}

	}
}
//todo 分发处理消息
func parseMsg(sendNode *Node, data []byte) {
	msg := new(utils.Msg)
	err := json.Unmarshal(data, msg)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	//存储时间戳
	msg.Date = time.Now().Unix()
	switch msg.Cmd {
		//发送数据
		case utils.CMD_MSG:
			//判断发送方向
			var subid int
			if msg.SrcType == utils.SRCTYPE_IP {
				subid = 0
			} else if msg.SrcType == utils.SRCTYPE_USER {
				subid = msg.IpId
			} else {
				return
			}
			//查找目标节点
			rwlocker.RLock()
			node, ok := GroupClientMap[msg.GroupID][subid]
			rwlocker.RUnlock()

			if msg.SrcType == utils.SRCTYPE_IP && ok {
				logs.Informational("類型1")
				//ip用户发出&&找到目标节点-》消息存储-》消息转发
				//-》消息存储
				if err := services.AddChatMsg(sendNode.SubId, msg.GroupID, msg.SrcType, msg.Data, models.READ_NO, time.Now().Unix()); err != nil {
					logs.Error(err.Error())
				}
				//存儲自身的iid
				msg.Iid = sendNode.SubId
				//-》消息转发
				resp, err := json.Marshal(msg)
				if err != nil {
					logs.Error(err.Error())
					return
				}
				node.Data <- resp
			} else if msg.SrcType == utils.SRCTYPE_IP && !ok {
				logs.Informational("類型2")

				//ip用户发出&&未找到节点-》消息存储-》通知邮件-》返回状态
				//-》消息存储
				if err := services.AddChatMsg(sendNode.SubId, msg.GroupID, msg.SrcType, msg.Data, models.READ_NO, time.Now().Unix()); err != nil {
					logs.Error(err.Error())
				}
				//-》通知邮件

				//-》返回状态
				restruct := utils.Msg{
					GroupID: msg.GroupID,
					IpId: sendNode.SubId,
					SrcType: utils.SRCTYPE_ORDER,
					Cmd: utils.CMD_NOTICE,
					Data: "",
				}
				resp, err := json.Marshal(restruct)
				if err != nil {
					logs.Error(err.Error())
					return
				}
				sendNode.Data <- resp

			} else if msg.SrcType == utils.SRCTYPE_USER && ok{
				logs.Informational("類型3")
				//后台用户发出&&找到目标节点-》消息存储-》消息转发
				//-》消息存储
				if err := services.AddChatMsg(msg.IpId, msg.GroupID, msg.SrcType, msg.Data, models.READ_NO, time.Now().Unix()); err != nil {
					logs.Error(err.Error())
				}
				//-》消息转发
				resp, err := json.Marshal(msg)
				if err != nil {
					logs.Error(err.Error())
					return
				}
				node.Data <- resp
			} else if msg.SrcType == utils.SRCTYPE_USER && !ok {
				logs.Informational("類型4")
				//后台用户发出&&未找到节点-》返回状态
				//-》返回状态
				restruct := utils.Msg{
					GroupID: msg.GroupID,
					IpId: msg.IpId,
					SrcType: utils.SRCTYPE_ORDER,
					Cmd: utils.CMD_NOTICE,
					Data: "",
				}
				resp, err := json.Marshal(restruct)
				if err != nil {
					logs.Error(err.Error())
					return
				}
				sendNode.Data <- resp
			}

			return
		//提交表单
		case utils.CMD_FORM:
			return
		case utils.CMD_NOTICE:
			return
	}

}
//todo 删除关系映射表的节点信息
func deleteNode(groupid, subid int) error {
	rwlocker.Lock()
	delete(GroupClientMap[groupid], subid)
	rwlocker.Unlock()

	if _, ok := GroupClientMap[groupid][subid]; ok {
		return errors.New("删除失败")
	}

	return nil
}
//todo 获取当前ip
func Ip(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.Header.Get("X-Real-IP"))
}
//todo 模拟发送消息
func Send(w http.ResponseWriter, r *http.Request) {
	rwlocker.RLock()
	node, ok := GroupClientMap[1][1]
	if !ok {
		logs.Error("没有找到目标节点")
		io.WriteString(w,"没有找到目标节点")
		return
	}
	rwlocker.RUnlock()
	msg := utils.Msg{
		GroupID: 1,
		IpId: 1,
		SrcType: "user",
		Cmd: "msg",
		Data: "模拟消息",
		Date: time.Now().Unix(),
	}
	send, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err.Error())
		io.WriteString(w, err.Error())
		return
	}
	node.Conn.WriteMessage(websocket.TextMessage, send)
}
//todo 加载聊天记录
func LoadChatList (w http.ResponseWriter, r *http.Request) {
	arg := utils.ChatListArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//获取聊天记录
	chats, err := services.ListChatMsg(arg.Iid, arg.Id)
	if err != nil {
		logs.Informational(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	utils.RespOkList(w, chats, len(chats))
}
//todo 将单个ipuser的所有聊天记录设置成已读
func SetChatRead(w http.ResponseWriter, r *http.Request) {
	arg := utils.ChatListArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//将聊天记录设置成已读
	if err := services.ChatReadHad(arg.Iid); err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	//返回成功
	utils.RespOk(w, nil, "", "")

}
//todo 打印关系映射表的信息
func DumpMap(w http.ResponseWriter, r *http.Request) {
	for k, v := range GroupClientMap {
		io.WriteString(w,"groupid:" +  strconv.Itoa(k))
		for i, _ := range v {
			io.WriteString(w, "--subid:" + strconv.Itoa(i))
		}
	}
}