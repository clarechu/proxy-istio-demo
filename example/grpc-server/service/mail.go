package service

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
)

const (
	me_mail string = "1062186165@qq.com"
)

var (
	from      = ""
	plainAuth smtp.Auth
)

func init() {
	from = os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")
	plainAuth = smtp.PlainAuth("", from, password, "smtp.qiye.aliyun.com")
}

// SendMailByServer 发送email 邮箱
func SendMailByServer() error {
	log.Println("start email ...")
	e := email.NewEmail()
	e.From = from
	e.To = []string{me_mail}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	err := e.SendWithTLS("smtp.qiye.aliyun.com:465", plainAuth,
		&tls.Config{
			ServerName:         "smtp.qiye.aliyun.com",
			InsecureSkipVerify: true})
	if err != nil {
		log.Printf("send fail ... \n")
		return err
	}
	log.Printf("send success ... \n")
	return err
}
