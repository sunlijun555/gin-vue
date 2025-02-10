package main

import (
	"fmt"
	"go-vue/config"
	"go-vue/core"
	"go-vue/router"

	_ "go-vue/docs"

	"github.com/sirupsen/logrus"
)

// @title 后台系统
// @version 1.0
// @description 后台API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println(config.Config.System)
	fmt.Println(config.Config.Logger)
	fmt.Println(config.Config.Mysql)
	fmt.Println(config.Config.Redis)
	core.InitDefaultLogger()
	logrus.Warnln("日志")
	logrus.Error("日志")
	logrus.Infof("日志")

	core.MysqlInit()
	core.RedisInit()
	router := router.InitRouter()
	address := fmt.Sprintf("%s:%d", config.Config.System.Host, config.Config.System.Port)
	logrus.Infof("启动系统，运行在：%s", address)
	router.Run(address)
}
