package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
   {
	"code":10001, // 程序中的错误
	"msg":xxx,    // 错误的提示信息
	"data":{},    // 返回的数据
   }
*/

// ResponseData 返回数据

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data,omitempty"` // omitempty 字段是空的时候会忽略
}

// ResponseError 返回错误的信息

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 返回错误信息，自定义错误信息

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess  返回成功的信息

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
