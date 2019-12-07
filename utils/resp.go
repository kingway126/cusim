package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type RespData struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg, omitempty"`
	Data  interface{} `json:"data, omitempty"`
	Rows  interface{} `json:"rows, omitempty"`
	Path   string      `json:"path, omitempty"`
	Total int         `json:"total, omitempty"`
}

//todo 响应失败函数
func RespFail(w http.ResponseWriter, msg string, url string) {
	Resp(w, -1, msg, nil, nil, 0, url)
}

//todo 响应成功函数
func RespOk(w http.ResponseWriter, data interface{}, msg string, url string) {
	Resp(w, 0, msg, data, nil, 0, url)
}

//todo 响应数据列表函数
func RespOkList(w http.ResponseWriter, rows interface{}, total int) {
	Resp(w, 0, "", nil, rows, total, "")
}


func Resp(w http.ResponseWriter, code int, msg string, data interface{}, rows interface{}, total int, url string) {
	//设置header为Json
	w.Header().Set("Content-Type", "application/json")
	//设置返回的状态
	w.WriteHeader(http.StatusOK)
	//定义消息结构体
	resp := RespData{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Rows:  rows,
		Path:   url,
		Total: total,
	}
	//将结构体转为json
	respJson, err := json.Marshal(resp)
	if err != nil {
		logs.Informational(err.Error())
		return
	}
	w.Write(respJson)
}
