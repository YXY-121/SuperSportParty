package main

import (
	"apiproject/client/process"
)

func main()  {
client:=process.NewClientProcess()
client.SendMessageToServer("1","2","发送信息","single")
	for  {
		
	}
}