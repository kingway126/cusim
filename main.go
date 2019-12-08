package main

import (
	"log"
	"net/http"
	"text/template"
	"CustomIM/controllers"
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

	//todo api注册
	http.HandleFunc("/api/login", controllers.LoginCheck) 				//登陆路由
	http.HandleFunc("/api/user/check", controllers.CheckToken)			//检验权限

	http.HandleFunc("/api/app/list", controllers.ListApp)				//获取多条app信息
	http.HandleFunc("/api/app/delete", controllers.DeleteApp)			//删除app信息
	http.HandleFunc("/api/app/resetuuid",controllers.UpdateAppUUID)	//更新app的uuid
	http.HandleFunc("/api/app/add", controllers.NewApp)					//插入新app
	http.HandleFunc("/api/app/update", controllers.ChangeApp)			//更新app
	http.HandleFunc("/api/app",controllers.GetApp)						//获取app信息

	http.HandleFunc("/api/chat",controllers.Chat)						//通讯
	http.HandleFunc("/api/chat/send", controllers.Send)

	//todo 文件服务注册
	http.Handle("/assets/", http.FileServer(http.Dir(".")))

	//todo 开启服务
	http.ListenAndServe(":3000", nil)
}
