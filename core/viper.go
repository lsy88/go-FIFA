package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Mode         string `mapstructure:"mode"`
	Port         string `mapstructure:"port"`
	URL          string `mapstructure:"url"`
	*PgSQLConfig `mapstructure:"pgsql"`
	*LogConfig   `mapstructure:"log"`
	*Jwt         `mapstructure:"jwt"`
}

type Jwt struct {
	SigningKey string `mapstructure:"signingkey"`
}

type PgSQLConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func ViperInit() error {
	//指定配置文件路径
	viper.SetConfigFile("./config/config.yaml")
	// 监控配置文件变化
	viper.WatchConfig()
	// 配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	//查找并读取配置文件的错误
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	//反序列化  将新配置反序列化到我们运行时的配置结构体中
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
