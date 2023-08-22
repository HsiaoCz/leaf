package email

import (
	"log"
	"net/smtp"
)

// 封装一个发送邮件的方法
func SendEmail(c *ConfigInfo, e EmailContent) error {
	//拼接smtp服务器地址
	smtpAddr := c.SmtpAddr + ":" + c.SmtpPort
	//认证信息
	auth := smtp.PlainAuth("", e.FromAddr, c.Secret, c.SmtpAddr)
	//配置邮件内容
	if e.ContentType == "html" {
		e.ContentType = "Content-Type:text/html;charset=UTF-8"
	} else {
		e.ContentType = "Content-Type:text/plain;charset=UTF-8"
	}
	//当有多个收件人的时候
	for _, to := range e.ToAddr {
		msg := []byte("To:" + to + "\r\n" + "From: " + e.FromAddr + "\r\n" + "Subject:" + e.Theme + "\r\n" + e.ContentType + "\r\n\r\n" + "<html><h1>" + e.Message + "</h1></html>")
		err := smtp.SendMail(smtpAddr, auth, e.FromAddr, []string{to}, msg)
		if err != nil {
			log.Fatalln("send mail err:", err)
		}
	}
	return nil
}
