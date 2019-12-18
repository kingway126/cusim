package utils

import (
	"gopkg.in/gomail.v2"
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
)
//控制发送email协程的数量
type Mail struct {
	Work 		int
	Run 		int
	Eve 		int
	Live 		time.Duration
	MailMsg		chan *gomail.Message
}

const (
	HOST 	= "smtp.exmail.qq.com"
	NAME 	= "recardo@supernovachina.com"
	PORT 	= 465
	SECRET = "PEV9cg5vry5cNYsf"
)


var mailAuto *Mail

//todo 初始化邮件发送功能
func MailInit()  {

	//初始化Mail
	mailAuto = new(Mail)
	mailAuto.Work = 0
	mailAuto.Run = 0
	mailAuto.Eve = 10
	mailAuto.Live = 30
	mailAuto.MailMsg = make(chan *gomail.Message, 100)
}
//todo 打印工作情况
func MailDump() {
	fmt.Println("work:", mailAuto.Work, "||run:", mailAuto.Run)
}
//todo 新增任务
func SendMsg(tit, msg string, email ...string) {
	m := gomail.NewMessage()
	m.SetHeader("From", NAME)
	m.SetHeader("To", email...)

	m.SetHeader("Subject", tit)
	m.SetBody("text/html", msg)
	//控制协程
	mailAuto.Work++
	mailAuto.Balance()
	//发送消息
	mailAuto.MailMsg <- m
}
//todo 新增协程
func (c *Mail) AddRun() {
	c.Run++
	go c.sendMail()
}
//todo 控制协程数量
func (c *Mail) Balance() {
	need, num := c.Work / c.Eve, c.Work % c.Eve
	if num > 0 {
		need++
	}
	if need > c.Run {
		//创建协程
		for i := 0; i < need - c.Run; i++ {
			c.AddRun()
		}
	}
}
// todo 发送邮件经常，30秒内自动退出
func (c *Mail) sendMail() {
	d := gomail.NewDialer(HOST, PORT, NAME, SECRET)

	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
			case m, ok := <-c.MailMsg:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						logs.Error(err.Error())
						return
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					logs.Error(err.Error())
				}
				//发送完消息，work减少1
				c.Work--

				// Close the connection to the SMTP server if no email was sent in
				// the last 30 seconds.
			case <-time.After(c.Live * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
				c.Run--
				//退出协程
				return
			}
	}
}