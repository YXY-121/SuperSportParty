package process

import (
	"apiproject/client/utils"
	"apiproject/common"
	"apiproject/config"
	"encoding/json"
	"net"
)

type ClientProcess struct {

}

func NewClientProcess() *ClientProcess {
	return &ClientProcess{}
}
func (c *ClientProcess)SendMessageToServer(senderId string,accepterId string,content  string, Type string)  {
	// connect server
	serverInfo := config.Configuration.ServerInfo
	conn, err := net.Dial("tcp", serverInfo.Host)

	if err != nil {
		return
	}

	var message common.Message
	message.Type = Type

	// group message
	singleMessage := common.SingleMessage{
		SenderId:  senderId,
		AccepterId: accepterId,
		Data:  content,
	}
	data, err := json.Marshal(singleMessage)
	if err != nil {
		return
	}

	message.Data = string(data)
	data, _ = json.Marshal(message)

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)

	return
}