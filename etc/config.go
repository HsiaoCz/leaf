package etc

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(LeafConf)

type LeafConf struct {
	App   AppConf   `mapstructure:"app"`
	Log   LogConf   `mapstructure:"log"`
	Mysql MysqlConf `mapstructure:"mysql"`
	Redis RedisConf `mapstructure:"redis"`
	Mongo MongoConf `mapstructure:"mongo"`
}

type AppConf struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"dev"`
	AppPort   string `mapstructure:"app_port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int    `mapstructure:"machine_id"`
}

type LogConf struct {
	Level      string `mapstructure:"debug"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConf struct {
	MysqlHost     string `mapstructure:"mysql_host"`
	MysqlPort     int    `mapstructure:"mysql_port"`
	User          string `mapstructure:"user"`
	MysqlPassword string `mapstructure:"mysql_password"`
	Dbname        string `mapstructure:"dbname"`
	MaxOpenConns  string `mapstructure:"mx_open_conns"`
	MaxIdleConns  string `mapstructure:"max_idle_conns"`
}

type RedisConf struct {
	RedisHost     string `mapstructure:"redis_host"`
	RedisPort     int    `mapstructure:"redis_port"`
	DB            int    `mapstructure:"db"`
	RedisPassword string `mapstructure:"redis_password"`
	PoolSize      int    `mapstructure:"poolsize"`
}

type MongoConf struct {
	MongoHost string `mapstructure:"mongo_host"`
	MongoPort int    `mapstructure:"mongo_port"`
	DbName    string `mapstructure:"db_name"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
}

func Init() (err error) {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("viper read config err:%v\n", zap.Error(err))
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper unmarshal err:%v\n", zap.Error(err))
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("config changed....")
		if err := viper.Unmarshal(Conf); err != nil {
			return
		}
	})
	return err
}
