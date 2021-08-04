package controllers

import (
	"apiproject/common"
	"apiproject/config"
	myutils2 "apiproject/server/utils"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"net"
)

type ClientController struct {
	beego.Controller
}
//type SingleMessage struct {
//	SenderId string
//	Data string
//	Type string
//	AccepterId string
//}
func (i* ClientController)Post()   {
	//在这里包装出提交的信息
	//先模拟单独消息
	data:=i.Ctx.Input.RequestBody
	fmt.Println(data)
	serverInfo:=config.Configuration.ServerInfo
	conn,err:=net.Dial("tcp",serverInfo.Host)
	if err!=nil {

	}
	message:=common.Message{}
	message.Data=string(data)
	message.Type="single"
	byteMessage,_:=json.Marshal(message)

	myutils2.DealMessage{Conn: conn}.SendDataToserver(byteMessage)

}

