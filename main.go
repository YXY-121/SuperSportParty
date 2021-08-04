package main

import (
	"apiproject/config"
	_ "apiproject/routers"
	"apiproject/server/process"
	"fmt"
	"net"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.Run()//这是http
	// 下面是tcp连接

	serverInfo := config.Configuration.ServerInfo
	fmt.Println("serverInfo", serverInfo)
	listener,err:=net.Listen("tcp",serverInfo.Host)
	defer listener.Close()
	if err!=nil {
		fmt.Printf("some error when run server, error: %v", err)
	}

	//永久跑着，等待客户端的连接
	for  {
		fmt.Println("等人中")
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Printf("some error when run server, error: %v", err)
		}
		//启动协程
		go dialogue(conn)

	}

}

func dialogue(conn net.Conn)  {
	defer conn.Close()
	connService:=process.ConService{conn}
	connService.DealConFromClient()
}

