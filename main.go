package main

import (
	"github/HsiaoCz/leaf/etc"
	"github/HsiaoCz/leaf/logger"
	"github/HsiaoCz/leaf/router"

	"go.uber.org/zap"
)

func main() {
	if err := etc.Init(); err != nil {
		zap.L().Error("etc init err:%v\n", zap.Error(err))
		return
	}
	if err := logger.InitLogger(logger.NewZapLoggerConf()); err != nil {
		zap.L().Error("Init logger err:%v\n", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	if err := router.Router(etc.Conf.App.Mode, etc.Conf.App.AppPort); err != nil {
		zap.L().Error("http server start fialed:%v\n", zap.Error(err))
		return
	}
}
