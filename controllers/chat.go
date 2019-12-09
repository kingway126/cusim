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
	Aid 		int
}
//关系映射表
var GroupClientMap  = make(map[int]map[int]*Node, 0)
//读写锁
var rwlocker sync.RWMutex
//todo websocket通讯接口
func Chat(w http.ResponseWriter, r *http.Request) {
	//获取请求的参数
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
	//检验接入是否合法
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
		logs.Info("ip:", ip, "remoteaddr:", r.RemoteAddr)
		//判断是否存在Ip用户，没有则创建
		ipuser, err = services.FindOrCreateIpUser(app.Id, app.Uid, ip)
		if err != nil {
			logs.Error(err.Error())
			return
		}
	}
	//创建连接节点
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
	//将键名与节点进行绑定
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
		//存储aid
		node.Aid = app.Id
		node.GroupId = ipuser.Uid
		node.SubId = ipuser.Id
		if GroupClientMap[ipuser.Uid] == nil {
			GroupClientMap[ipuser.Uid] = make(map[int]*Node)
		}
		rwlocker.Lock()
		GroupClientMap[ipuser.Uid][ipuser.Id] = node
		rwlocker.Unlock()
	}
	//开启监听websocket协程
	go ListenWebSocket(node)
	//开启监听NODE协程
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
		parseMsg(node.SubId,[]byte(msg))
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
func parseMsg(id int, data []byte) {
	msg := new(utils.Msg)
	err := json.Unmarshal(data, msg)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	switch msg.Cmd {
		//发送数据
		case utils.CMD_MSG:
			msg.SrcId = id
			//查找目标节点
			rwlocker.RLock()
			node := GroupClientMap[msg.DstGroup][msg.DstId]
			rwlocker.RUnlock()
			//消息持久化
			if msg.SrcType == utils.SRCTYPE_IP {
				err := services.AddChatMsg(node.Aid, node.SubId, node.GroupId,msg.SrcType, msg.Data, models.READ_NO,time.Now().Unix())
				if err != nil {
					logs.Error(err.Error())
				}
			} else if msg.SrcType == utils.SRCTYPE_USER {
				//获取aid
				ipuser, err := services.GetIpUserById(msg.DstId)
				if err != nil {
					logs.Error(err.Error())
				} else if ipuser != nil {
					err = services.AddChatMsg(ipuser.Aid, msg.DstId, node.GroupId, msg.SrcType, msg.Data, models.READ_YES, time.Now().Unix())
					if err != nil {
						logs.Error(err.Error())
					}
				}

				//发送邮件通知
			}
			//处理消息
			resp, err := json.Marshal(msg)
			if err != nil {
				logs.Error(err.Error())
				return
			}
			//发送数据
			node.Data <- resp
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
	io.WriteString(w, r.Header.Get("remoteaddr"))
	io.WriteString(w, utils.ParseAddr(r.RemoteAddr))
}
//todo 模拟发送消息
func Send(w http.ResponseWriter, r *http.Request) {
	rwlocker.RLock()
	node := GroupClientMap[1][1]
	rwlocker.RUnlock()
	msg := utils.Msg{
		DstGroup: 1,
		DstId: 1,
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
