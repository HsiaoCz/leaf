package router

import (
	"github/HsiaoCz/leaf/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(r *fiber.App) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				// 用户注册
				auth.Post("/register", controller.UserRegister)
				// 用户登录
				auth.Post("/login", controller.UserLogin)
			}
		}
	}
}
