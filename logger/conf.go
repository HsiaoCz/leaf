package logger

import "github/HsiaoCz/leaf/etc"

type ZapLoggerConf struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LogLevel   string
	AppMode    string
}

func NewZapLoggerConf() ZapLoggerConf {
	return ZapLoggerConf{
		Filename:   etc.Conf.Log.Filename,
		MaxSize:    etc.Conf.Log.MaxSize,
		MaxBackups: etc.Conf.Log.MaxBackups,
		MaxAge:     etc.Conf.Log.MaxAge,
		LogLevel:   etc.Conf.Log.Level,
		AppMode:    etc.Conf.App.Mode,
	}
}
