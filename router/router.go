package router

import (
	"github/HsiaoCz/leaf/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				// 用户注册
				auth.POST("/register", controller.UserRegister)
				// 用户登录
				auth.POST("/login", controller.UserLogin)
			}
		}
	}
}
