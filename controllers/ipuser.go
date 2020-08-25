package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/recardoz/cusim/services"
	"github.com/recardoz/cusim/utils"
	"net/http"
	"time"
)

//todo 获取ipuser列表
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
	begin := end - 3600*24*3
	ipusers, err := services.ListIpUserForTime(arg.Id, begin, end)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	utils.RespOkList(w, ipusers, len(ipusers))
}

//todo 获取单个的ipuser
func GetIpUser(w http.ResponseWriter, r *http.Request) {
	arg := utils.ChatListArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		utils.RespFail(w, "请求参数错误", "")
		logs.Error(err.Error())
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//获取ipuser信息
	ipuser, err := services.GetIpUser(arg.Iid, arg.Id)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, "获取记录失败", "")
		return
	}
	utils.RespOk(w, ipuser, "", "")

}
