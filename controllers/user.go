package controllers

import (
	"net/http"
	"CustomIM/services"
	"CustomIM/utils"
	"github.com/astaxie/beego/logs"
)

//todo 登陆api
func LoginCheck(w http.ResponseWriter, r *http.Request) {
	//解析请求参数
	arg := utils.LoginArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Informational(err.Error())
		utils.RespFail(w, "请求数据格式错误！", "")
		return
	}

	//判断账号是否正确
	user, err := services.GetUserInfo(arg.User)
	if err != nil {
		logs.Informational(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	sha1Pass := utils.Sha1Pwd(arg.Pass)
	if user.Pwd != sha1Pass {
		logs.Informational("账号或者密码错误")
		utils.RespFail(w, "账号或者密码错误", "")
		return
	}

	//更新token
	token, err := services.UpdateUserToken(user.User)
	if err != nil {
		logs.Informational(err.Error())
		utils.RespFail(w, "数据处理失败，请联系后台管理员", "")
		return
	}

	//检测UUID
	if err := services.CheckUUID(arg.User); err != nil {
		logs.Informational(err.Error())
	}

	//返回token信息
	resp := utils.TokenArgs{
		Id:    user.Id,
		Token: token,
	}
	utils.RespOk(w, resp, "", "/index")
}

//todo 鉴权api
func CheckToken(w http.ResponseWriter, r *http.Request) {

	arg := utils.TokenArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Informational(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	user, err := services.GetUserInfoById(arg.Id)
	if err != nil {
		logs.Informational(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	if user.Hash != arg.Token {
		utils.RespFail(w, "该用户未登录", "")
		return
	}

	utils.RespOk(w, user.Uuid, "该用户已经登陆","")

}
//todo 检测用户token
func TokenIsRight(id int, token string) bool {
	user, err := services.GetUserInfoById(id)
	if err != nil {
		logs.Informational(err.Error())
		return false
	} else if user.Hash != token {
		return false
	}
	return true
}