package main

import (
	"ys_go/global"
	"ys_go/initialize/config"
	"ys_go/initialize/gorm"
	"ys_go/initialize/logger"
	"ys_go/initialize/router"
)

func main() {
	// 1. 获取配置信息
	global.Config = config.InitConfig()

	// 2. 初始化日志
	log, err := logger.NewLogger("development", "logs", "[myApp] ")
	if err != nil {
		panic(err)
	}
	global.Log = log
	defer log.Close()

	// 3. 初始化gorm
	global.DB = gorm.InitDB()

	// 4. gin
	router.InitRouter()

}
