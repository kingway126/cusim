package controllers

import (
	"net/http"
	"CustomIM/utils"
	"github.com/astaxie/beego/logs"
	"time"
	"CustomIM/services"
)

func GetIpUserList(w http.ResponseWriter, r *http.Request) {
	arg := utils.TokenArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Error("err0")
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//获取最近3天链接的用户列表
	end := time.Now().Unix()
	begin := end - 3600 * 24 * 3
	ipusers, err := services.ListIpUserForTime(arg.Id, begin, end)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	utils.RespOkList(w, ipusers, len(ipusers))
}
