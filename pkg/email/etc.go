package email

//邮件相关的model
// 邮箱配置服务

type ConfigInfo struct {
	SmtpAddr string
	SmtpPort string
	Secret   string
}

// 邮件内容

type EmailContent struct {
	FromAddr    string
	ContentType string
	Theme       string
	Message     string
	ToAddr      []string
}

//验证码的过期时间

var CodeExpire = 600
