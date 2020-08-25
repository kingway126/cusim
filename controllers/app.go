package controllers

import (
	"github.com/recardoz/cusim/services"
	"github.com/recardoz/cusim/utils"
	"net/http"
)

//todo 获取站点信息列表 api
func ListApp(w http.ResponseWriter, r *http.Request) {
	arg := new(utils.AppArgs)
	if err := utils.Bind(r, arg); err != nil {
		utils.RespFail(w, "请求参数错误", "")
		return
	}

	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}

	all, apps, err := services.ListAppInfo(arg.Id, arg.PageIndex, arg.PageSize, arg.Search)
	if err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}

	utils.RespOkList(w, apps, all)
}

//todo 删除站点 api
func DeleteApp(w http.ResponseWriter, r *http.Request) {
	arg := new(utils.DeleteSiteArgs)
	if err := utils.Bind(r, arg); err != nil {
		utils.RespFail(w, "请求参数错误", "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//删除数据
	if err := services.DeleteApp(arg.Aid); err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}
	//删除数据成功
	utils.RespOk(w, nil, "删除站点成功", "")
}

//todo 重置站点UUID api
func UpdateAppUUID(w http.ResponseWriter, r *http.Request) {
	arg := new(utils.DeleteSiteArgs)
	if err := utils.Bind(r, arg); err != nil {
		utils.RespFail(w, "请求参数错误", "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//重置uuid
	if err := services.ResetAppUuid(arg.Aid); err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}
	//返回成功
	utils.RespOk(w, nil, "重置UUID成功", "")
}

//todo 插入新站点
func NewApp(w http.ResponseWriter, r *http.Request) {
	arg := new(utils.SiteArgs)
	if err := utils.Bind(r, arg); err != nil {
		utils.RespFail(w, "请求参数错误", "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//添加新数据
	if err := services.NewApp(arg.Id, arg.Name, arg.Url, ""); err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}
	//返回成功
	utils.RespOk(w, nil, "添加站点成功", "/sitemap")
}

//todo 修改站点信息
func ChangeApp(w http.ResponseWriter, r *http.Request) {
	arg := new(utils.SiteArgs)
	if err := utils.Bind(r, arg); err != nil {
		utils.RespFail(w, "请求参数错误", "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//修改数据
	if err := services.ChangeApp(arg.Aid, arg.Name, arg.Url, ""); err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}
	//返回成
	utils.RespOk(w, nil, "修改站点成功", "/sitemap")
}

//todo 获取一条站点信息
func GetApp(w http.ResponseWriter, r *http.Request) {
	arg := new(utils.DeleteSiteArgs)
	if err := utils.Bind(r, arg); err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}
	//检验权限
	if ok := TokenIsRight(arg.Id, arg.Token); !ok {
		utils.RespFail(w, "登陆过期，请重新登陆", "/user/login")
		return
	}
	//获取站点信息
	app, err := services.GetAppInfo(arg.Aid)
	if err != nil {
		utils.RespFail(w, err.Error(), "")
		return
	}
	//返回数据
	utils.RespOk(w, app, "", "")
}
