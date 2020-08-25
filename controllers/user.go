package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/recardoz/cusim/models"
	"github.com/recardoz/cusim/services"
	"github.com/recardoz/cusim/utils"
	"net/http"
)

//todo 登陆api
func LoginCheck(w http.ResponseWriter, r *http.Request) {
	//解析请求参数
	arg := utils.LoginArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, "请求数据格式错误！", "")
		return
	}

	//判断账号是否正确
	user, err := services.GetUserInfo(arg.User)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	sha1Pass := utils.Sha1Pwd(arg.Pass)
	if user.Pwd != sha1Pass {
		logs.Error("账号或者密码错误")
		utils.RespFail(w, "账号或者密码错误", "")
		return
	}

	//更新token
	token, err := services.UpdateUserToken(user.User)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, "数据处理失败，请联系后台管理员", "")
		return
	}

	//检测UUID
	uuid, err := services.CheckUUID(arg.User)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, "系统出错，请联系管理员", "")
	}

	//返回token信息
	resp := utils.LoginRespArgs{
		Id:    user.Id,
		Token: token,
		Uuid:  uuid,
	}
	utils.RespOk(w, resp, "", "/index")
}

//todo 鉴权api
func CheckToken(w http.ResponseWriter, r *http.Request) {

	arg := utils.TokenArgs{}
	if err := utils.Bind(r, &arg); err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	user, err := services.GetUserInfoById(arg.Id)
	if err != nil {
		logs.Error(err.Error())
		utils.RespFail(w, err.Error(), "")
		return
	}
	if user.Hash != arg.Token {
		utils.RespFail(w, "该用户未登录", "")
		return
	}

	utils.RespOk(w, user.Uuid, "该用户已经登陆", "")

}

//todo 检测用户token
func TokenIsRight(id int, token string) bool {
	user, err := services.GetUserInfoById(id)
	if err != nil {
		logs.Error(err.Error())
		return false
	} else if user.Hash != token {
		return false
	}
	return true
}

//todo 获取用户信息
func UserInfo(w http.ResponseWriter, r *http.Request) {
	arg := utils.TokenArgs{}
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
	//获取用户的信息
	user, err := services.GetUserInfoById(arg.Id)
	if err != nil {
		utils.RespFail(w, "获取数据失败", "")
		logs.Error(err.Error())
		return
	}
	utils.RespOk(w, user, "", "")
}

//todo 更新用户邮箱
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	arg := utils.UserInfo{}
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
	//更新用户邮箱
	if err := services.UpdateEmail(arg.Id, arg.Email); err != nil {
		utils.RespFail(w, "更新账户失败", "")
		logs.Error(err.Error())
		return
	}

	utils.RespOk(w, nil, "更新账户成功", "")
}

//todo 更新用户密码
func UpdateUserPwd(w http.ResponseWriter, r *http.Request) {
	arg := utils.UserPwd{}
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
	//加密密码
	pwd := utils.Sha1Pwd(arg.Pwd)
	//更新密码
	if err := services.UpdatePwd(arg.Id, pwd); err != nil {
		utils.RespFail(w, "修改密码失败", "")
		logs.Error(err.Error())
		return
	}

	utils.RespOk(w, nil, "修改密码成功", "/pwd")
}

//todo 获取首页用户的信息
func GetIndexNum(w http.ResponseWriter, r *http.Request) {
	arg := utils.TokenArgs{}
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

	index := new(utils.IndexNum)
	var err error
	//查app数量
	if index.App, err = services.GetAllApp(arg.Id); err != nil {
		utils.RespFail(w, "获取数据失败", "")
		logs.Error(err.Error())
		return
	}
	//查ipuser数量
	if index.User, err = services.GetIpUserNum(arg.Id); err != nil {
		utils.RespFail(w, "获取数据失败", "")
		logs.Error(err.Error())
		return
	}
	//查未读消息数量
	if index.NoRead, err = services.GetNoReadNum(arg.Id, models.READ_NO); err != nil {
		utils.RespFail(w, "获取数据失败", "")
		logs.Error(err.Error())
		return
	}
	//查所有消息数量
	if index.Read, err = services.GetAllNum(arg.Id); err != nil {
		utils.RespFail(w, "获取数据失败", "")
		logs.Error(err.Error())
		return
	}

	utils.RespOk(w, index, "", "")
}
