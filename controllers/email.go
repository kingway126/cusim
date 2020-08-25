package controllers

import (
	"github.com/recardoz/cusim/utils"
	"net/http"
)

//todo 模拟发送邮件
func SendEmail(w http.ResponseWriter, r *http.Request) {
	utils.SendMsg("你好", "测试", "1098977435@qq.com")
	utils.MailDump()
}
