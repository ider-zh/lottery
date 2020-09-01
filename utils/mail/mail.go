package mail

import (
	"log"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func SendMail() {
	e := email.NewEmail()
	e.From = "Golang Loger <ruangnazi@163.com>"
	e.To = []string{"admin@knogen.cn"}
	// e.Bcc = []string{"752794935@qq.com"}
	// e.Cc = []string{"ruangnazi@gmail.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "ruangnazi@163.com", "001ider", "smtp.163.com"))
	if err != nil {
		log.Println(err)
	}
}

var EmailChan chan *email.Email

func init() {
	EmailChan = make(chan *email.Email)

	p, _ := email.NewPool(
		"smtp.163.com:25",
		1,
		smtp.PlainAuth("", "ruangnazi@163.com", "001ider", "smtp.163.com"),
	)

	go func() {
		for e := range EmailChan {
			// 失败的任务会重新进行，直到成功
			for {
				err := p.Send(e, 30*time.Second)
				time.Sleep(30 * time.Second)
				if err != nil {
					log.Println("mail send fail,", err)
				} else {
					break
				}
			}
		}
	}()
}

func NewSimpleTextMail(subject string, text string) {

	e := email.NewEmail()
	e.From = "SSQ 风云 <ruangnazi@163.com>"
	e.To = []string{"admin@knogen.cn"}
	// e.Cc = []string{"326737833@qq.com"}
	// e.Bcc = []string{"752794935@qq.com"}
	// e.Cc = []string{"lu.hongfei@shuzhi.ai"}
	e.Subject = subject
	e.Text = []byte(text)
	// e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	EmailChan <- e
}
