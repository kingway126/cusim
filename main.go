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

	//todo 路由注册
	http.HandleFunc("/api/login", controllers.LoginCheck) //登陆路由
	http.HandleFunc("/api/user/check", controllers.CheckToken)

	//todo 文件服务注册
	http.Handle("/assets/", http.FileServer(http.Dir(".")))

	//todo 开启服务
	http.ListenAndServe(":3000", nil)
}
