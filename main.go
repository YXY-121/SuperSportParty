package main

import (
	"apiproject/common"
	"apiproject/config"
	"apiproject/log"
	"apiproject/routers"
	service "apiproject/service/websocket"

	"github.com/sirupsen/logrus"
)

func main() {
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }

	// beego.Run()

	//初始化日志

	//初始化配置
	config.ConfigSetup()
	common.InitDB()
	log.SetupLog()
	//初始化群
	service.InitAllGroup()
	// //初始化路由
	err := routers.InitRouter().Run()
	if err != nil {
		logrus.Errorf("router init fail")
	}
	logrus.Println("完成初始化")

}
