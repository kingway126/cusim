package main

import (
	"log"
	"net/http"
	"text/template"
	"CustomIM/controllers"
	"CustomIM/utils"
)

//模板自动注册函数
func RegisterView() {
	//解析
	tpl, err := template.ParseGlob("views/**/*")
	if err != nil {
		//打印并直接退出
		log.Fatal(err.Error())

	}

	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, tplName, nil)
		})
	}
}

func main() {
	//todo 模板注册
	RegisterView()

	//todo 启动邮件服务
	utils.MailInit()

	//todo api注册
	http.HandleFunc("/api/login", controllers.LoginCheck) 				//登陆路由
	http.HandleFunc("/api/index", controllers.GetIndexNum)				//获取首页的信息

	http.HandleFunc("/api/user/check", controllers.CheckToken)			//检验权限
	http.HandleFunc("/api/user", controllers.UserInfo)					//获取用户信息
	http.HandleFunc("/api/user/email", controllers.UpdateUser)			//更新用户邮箱
	http.HandleFunc("/api/user/pwd", controllers.UpdateUserPwd)			//更新用户密码

	http.HandleFunc("/api/app/list", controllers.ListApp)				//获取多条app信息
	http.HandleFunc("/api/app/delete", controllers.DeleteApp)			//删除app信息
	http.HandleFunc("/api/app/resetuuid",controllers.UpdateAppUUID)	//更新app的uuid
	http.HandleFunc("/api/app/add", controllers.NewApp)					//插入新app
	http.HandleFunc("/api/app/update", controllers.ChangeApp)			//更新app
	http.HandleFunc("/api/app",controllers.GetApp)						//获取app信息

	http.HandleFunc("/api/chat",controllers.Chat)						//通讯
	http.HandleFunc("/api/chat/send", controllers.Send)					//模拟发送消息
	http.HandleFunc("/api/chat/ip", controllers.Ip)						//获取请求的IP地址
	http.HandleFunc("/api/chat/list", controllers.LoadChatList)			//获取聊天记录
	http.HandleFunc("/api/chat/read", controllers.SetChatRead)			//将聊天记录设置成已读
	http.HandleFunc("/api/chats/read", controllers.SetUserChatRead)		//将所有的聊天记录设置成已读
	http.HandleFunc("/api/chat/noread", controllers.GetNoReadNum)		//获取未读消息的条数

	//http.HandleFunc("/api/chat/dump", controllers.DumpMap)				//打印关系映射表

	http.HandleFunc("/api/ipuser/list", controllers.GetIpUserList)		//获取ipuser用户最近3天链接的用户
	http.HandleFunc("/api/ipuser", controllers.GetIpUser)				//获取单个ipuser信息

	http.HandleFunc("/api/email/send", controllers.SendEmail)			//模拟发送邮件


	//todo 文件服务注册
	http.Handle("/assets/", http.FileServer(http.Dir(".")))

	//todo 开启服务
	http.ListenAndServe(":2222", nil)
}
