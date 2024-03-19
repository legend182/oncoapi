package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name	   string `mapstructure:"name"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Mode	   string `mapstructure:"mode"`
	*MongodbConf
	*LogConf
}

type MongodbConf struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	DB         string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
}

type LogConf struct{
	Level       string      `mapstructure:"level"`
	Filename    string		`mapstructure:"filename"`
	Maxsize	    int			`mapstructure:"maxsize"`
	MaxAge		int			`mapstructure:"maxage"`
	MaxBackups	int			`mapstructure:"maxbackups"`
}

func Init() error {
	// 读取文件路径
	viper.SetConfigFile("./conf/conf.yaml")
	// 读取环境变量
	viper.WatchConfig()
	// 监听文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("修改配置文件")
		err := viper.Unmarshal(&Conf)
		if err!=nil{
			fmt.Println("配置更新成功")
		}
	})
	// 读取配置文件
	err:=viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err:=viper.Unmarshal(&Conf);err!=nil{
		panic(err)
	}
	return err
}