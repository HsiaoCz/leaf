package router

import (
	"github/HsiaoCz/leaf/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	// 用户模块
	user := r.Group("/api/v1/user")
	{
		// 用户注册
		user.POST("/register", controller.UserRegister)
		// 用户登录
		user.POST("/login", controller.UserLogin)
	}
}
