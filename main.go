package main

import (
	_ "apiproject/routers"
	webSocketServer "apiproject/websocket/server"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.Run()//这是http
	// 下面是tcp连接

	webSocketServer.WebServer()
}