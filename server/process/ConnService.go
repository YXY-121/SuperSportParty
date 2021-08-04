package process

import (
	"apiproject/common"

	"apiproject/server/utils"
	"fmt"
	"net"
)

type ConService struct {
	Conn net.Conn
}

func NewConService(Conn net.Conn) *ConService {
	return &ConService{
		Conn: Conn,
	}
}
func (c *ConService)DealConFromClient()  {
	//长轮询读取客户的信息
	for  {
		d:= utils.Dispatcher{Conn: c.Conn}
		//读取data的时候出错

		message,err:=d.ReadData()
		fmt.Println(message)
		if err!=nil {
		//todo 报错
			fmt.Println("报错了",err)
			return
		}
		//根据type处理message
	c.DealMessageByType(message)

	}
}

func (c *ConService)DealMessageByType(message common.Message)  {
	if message.Type=="single" {
		singleService:=NewSingleService()
		singleService.SendSingleMessage(message.Data)
	}else {
		groupService:=NewGroupService()
		groupService.SendGroupService(message.Data)
	}
}