package main

import (
	"apiproject/server/process"
	"net"
)

type A struct {
	One string
	Two string
}
type B struct {
	One string
	Three string
	Four string
}


func main() {

//	groupMap:=make(map[string]chan string)
//	if beego.BConfig.RunMode == "dev" {
//		beego.BConfig.WebConfig.DirectoryIndex = true
//		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
//	}
//	a:=A{
//		One: "1hi" ,
//		Two: "shfd",
//	}
//	b:=B{
//
//	}
//
//
//
//	bytes,err:=json.Marshal(&a)
//	if err!=nil {
//		fmt.Println("json.Marshal失败")
//	}
//
//	fmt.Println(bytes)
//	s:=string(bytes)
//	fmt.Println(s)
//	json.Unmarshal(bytes[:],&b)
//	fmt.Println(b)
//	fmt.Println("len",len(bytes))
////	beego.Run()//这是http
//	//下面是tcp连接
//
//	serverInfo := config.Configuration.ServerInfo
//	fmt.Println("serverInfo", serverInfo)
//	listener,err:=net.Listen("tcp",serverInfo.Host)
//	defer listener.Close()
//	if err!=nil {
//		fmt.Printf("some error when run server, error: %v", err)
//	}
//
//	//永久跑着，等待客户端的连接
//	for  {
//		fmt.Println("等人中")
//		conn,err:=listener.Accept()
//		if err!=nil {
//			fmt.Printf("some error when run server, error: %v", err)
//		}
//		//启动协程
//		go dialogue(conn)
//	}

}

func dialogue(conn net.Conn)  {
	defer conn.Close()
	connService:=process.ConService{conn}
	connService.DealConFromClient()
}

