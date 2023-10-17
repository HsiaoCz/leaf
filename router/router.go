package router

import (
	"github/HsiaoCz/leaf/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(r *fiber.App) {

	// 用户模块
	user := r.Group("/api/v1/user")
	{
		// 用户注册
		user.Post("/register", controller.UserRegister)
		// 用户登录
		user.Post("/login", controller.UserLogin)
	}
}
