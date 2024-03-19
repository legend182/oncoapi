package main

import (
	"fmt"
	"oncoapi/dao/mongodb"
	"oncoapi/logger"
	"oncoapi/setting"
)

func main() {
	// 读取配置文件
	if err := setting.Init();err!=nil{
		fmt.Println("配置文件读取失败:",err)
		return
	}
	// 配置log
	if err := logger.Init(setting.Conf.LogConf,setting.Conf.Mode);err!=nil{
		fmt.Println("初始化日志失败:",err)
	}
	// 初始化mongodb
	if err := mongodb.Init(setting.Conf.MongodbConf);err!=nil{
		fmt.Println("初始化mongodb失败:",err)
	}
	defer mongodb.Close()
	// 注册路由
	

}