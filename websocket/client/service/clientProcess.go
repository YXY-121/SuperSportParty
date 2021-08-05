package client

import (
	"apiproject/common"
	"apiproject/websocket/reponsitory"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type ClientService struct {

	Url url.URL
	Conn  *websocket.Conn
	HubMap map[string]*Hub
	AcceptedMessages chan[]byte
	SenderId string
}


func (c *ClientService)ReturnAccectped(){
	defer c.Conn.Close()
	for  {
		select {
		//通过通道发现收到了消息，就会返回给自己的客户端页面 writeMessage
		case message,ok:=<-c.AcceptedMessages:

			if !ok {
				fmt.Println("client不ok" )
			}

			c.Conn.WriteMessage(websocket.TextMessage,message)

		}
	}

}

//定义一个有所有群的hubmap总群
func (c *ClientService)CreateGroup()  {
	//创建hub
}
func  (c *ClientService)AddGroup(GroupId  string)  {
	//c.HubMap[GroupId]=
}
func (c *ClientService)SendOther()  {

	for{
		_,messageData,err:=c.Conn.ReadMessage()
		//这个很重要 不然client退出的时候 server会报错
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message:=common.Message{}
		messageData1:=messageData
		json.Unmarshal(messageData1,&message)

		//初始化client的群频道		//同时让client注册到这个频道里
		if message.UserId!="" {
			hubs:=c.GetClientHubMap(message.UserId)
			c.HubMap=hubs

		}


		if (message!=common.Message{})&&message.Type==common.SingleMessageType {
			c.SendSingleMessage(messageData)
		} else if (message!=common.Message{})&&message.Type==common.GroupMessageType {

			c.SendGroupMessage(messageData)
		}else {
			c.comonMessage(messageData)
		}

	}
}
func(c *ClientService)comonMessage(messageData[]byte)  {

	for _,hub:=range c.HubMap {
		a:="group"+hub.HubId
		a=a+string(messageData[:])

		fmt.Println(string(messageData[:]))
		bytemessage:=[]byte(a)
		hub.GroupBroadCast<-bytemessage
	}
}

func (c *ClientService)SendGroupMessage(messageData[]byte)  {


		message:=common.GroupMessage{}

		json.Unmarshal(messageData,&message)

		groupMessage,_:=json.Marshal(message)
		groupId:=message.GroupId
		//检查该client是否关注了这个群
		hub:=c.HubMap[groupId]
	if hub!=nil {
		fmt.Println("已发送给hub处理")
		hub.GroupBroadCast<-groupMessage
	}else  {
		fmt.Println("查不到这个群")
	}
	return


}
func (c *ClientService)SendSingleMessage(messageData[]byte)  {
	message:=common.SingleMessage{}
	json.Unmarshal(messageData,&message)

}

func (c *ClientService)GetClientHubMap(userId string)map[string]*Hub {
	groups:=reponsitory.GetGroupsByUserId(userId)
	fmt.Println(groups)
	hubMap:=make(map[string]*Hub)
	for _,v:=range groups{
		hubMap[v.GroupId]= AllHub[v.GroupId]
		AllHub[v.GroupId].Clients[c]=true
	}

	return hubMap
}
func InitAllGroup()  {
	groups:=reponsitory.GetAllGroups();
	for _,group:=range groups {
		hub:= NewHub(group.GroupId)
		AllHub[group.GroupId]=hub
		go hub.Run()
	}
	fmt.Println(len(AllHub))
}



