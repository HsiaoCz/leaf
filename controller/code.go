package controller

// 这里定义一些返回的错误码

// ResCoe 定义一个状态码类型

type ResCode int64

// const 里面定义一些状态码

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeInvalidEmailCode
	CodeInvalidPhoneCode
	CodeServerBusy

	CodeInvalidAuth
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeUserExist:        "用户名已经存在",
	CodeUserNotExist:     "用户不存在",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeInvalidEmailCode: "邮件验证码不正确",
	CodeInvalidPhoneCode: "短信验证码不正确",
	CodeInvalidAuth:      "无效的登录Token",
	CodeServerBusy:       "服务器繁忙",
	CodeNeedLogin:        "当前请求需要登录",
}

// Msg 状态码获取自身信息的方法

func (rc ResCode) Msg() string {
	msg, ok := codeMsgMap[rc]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
