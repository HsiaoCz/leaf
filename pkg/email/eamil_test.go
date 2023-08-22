package email

import (
	"log"
	"net/smtp"
	"testing"
)

// 邮箱配置服务
type configInfo struct {
	smtpAddr string
	smtpPort string
	secret   string
}

// 邮件内容
type emailContent struct {
	fromAddr    string
	contentType string
	theme       string
	message     string
	toAddr      []string
}

func sendEmail(c *configInfo, e *emailContent) error {
	//拼接smtp服务器地址
	smtpAddr := c.smtpAddr + ":" + c.smtpPort
	//认证信息
	auth := smtp.PlainAuth("", e.fromAddr, c.secret, c.smtpAddr)
	//配置邮件内容
	if e.contentType == "html" {
		e.contentType = "Content-Type:text/html;charset=UTF-8"
	} else {
		e.contentType = "Content-Type:text/plain;charset=UTF-8"
	}
	//当有多个收件人的时候
	for _, to := range e.toAddr {
		msg := []byte("To:" + to + "\r\n" + "From: " + e.fromAddr + "\r\n" + "Subject:" + e.theme + "\r\n" + e.contentType + "\r\n\r\n" + "<html><h1>" + e.message + "</h1></html>")
		err := smtp.SendMail(smtpAddr, auth, e.fromAddr, []string{to}, msg)
		if err != nil {
			log.Fatalln("send mail err:", err)
		}
	}
	return nil
}

func TestMail(t *testing.T) {
	//测试发送邮件
	//1.收集配置信息
	config := configInfo{
		//smtp服务器地址
		smtpAddr: "smtp.126.com",
		//smtp服务器密钥
		secret: "NWQSVVLEIAYDRPQU",
		//smtp服务器端口
		smtpPort: "25",
	}

	//收集邮件内容
	content := emailContent{
		//发件人
		fromAddr: "hsiaocz@126.com",
		//收件人
		toAddr: []string{"1299720482@qq.com", "984274788@qq.com"},
		//邮件格式
		contentType: "html",
		//邮件主题
		theme: "验证码",
		//邮件内容
		message: "123456",
	}

	//发送邮件
	err := sendEmail(&config, &content)
	if err != nil {
		log.Fatalln("send message err:", err)
	} else {
		log.Fatalln("send success")
	}
}
