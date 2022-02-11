package service

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

const (
	me_mail string = "1062186165@qq.com"
)

var (
	from      = ""
	port      = 465
	host      = "smtp.qiye.aliyun.com"
	plainAuth smtp.Auth
)

func init() {
	from = os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")
	host = os.Getenv("HOST")
	if host == "" {
		host = "smtp.qiye.aliyun.com"
	}
	password = "N0P@$$w0rd"
	p, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		port = p
	}
	log.Printf("MAIL_FROM :%s", from)
	log.Printf("MAIL_PASSWORD :%s", password)
	log.Printf("HOST :%s", host)
	log.Printf("PORT :%d", port)
	log.Printf("STARTTLS :%s", os.Getenv("STARTTLS"))
	plainAuth = smtp.PlainAuth("", from, password, host)
}

// SendMailByServer 发送email 邮箱
func SendMailByServer(message string) (err error) {
	log.Println("start email ...")
	e := email.NewEmail()
	e.From = from
	e.To = []string{me_mail}
	e.Subject = "Awesome Subject"
	e.Text = []byte(message)
	e.Timeout = 1 * time.Minute
	if os.Getenv("STARTTLS") == "true" {
		err = e.SendWithStartTLS(fmt.Sprintf("%s:%d", host, port), plainAuth,
			&tls.Config{
				ServerName:         host,
				InsecureSkipVerify: true})
	} else {

		err = e.SendWithTLS(fmt.Sprintf("%s:%d", host, port), plainAuth,
			&tls.Config{
				ServerName:         host,
				InsecureSkipVerify: true})
	}
	if err != nil {
		log.Printf("[ERROR] err:%v", err)
		log.Printf("send fail ... \n")
		return err
	}
	log.Printf("send success ... \n")
	return err
}
