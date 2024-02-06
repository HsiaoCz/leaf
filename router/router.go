package router

import (
	"github/HsiaoCz/leaf/controller"
	"github/HsiaoCz/leaf/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(mode string, addr string) error {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
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
	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	return srv.ListenAndServe()
}
