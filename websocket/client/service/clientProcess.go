package client

import (
	"apiproject/common"
	"apiproject/websocket/model"
	"apiproject/websocket/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"net/url"
	"time"
)

var groupHistory=repository.GroupHistoryMessageRepository{}
var singleHistory=repository.SingleHistoryMessageRepository{}

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

func (c *ClientService)SendOther()  {

	for{

		messageType,messageData,err:=c.Conn.ReadMessage()
		//这个很重要 不然client退出的时候 server会报错
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//注销
		if messageType==websocket.CloseMessage {
			fmt.Println("注销")
			c.LogOut()
		}

		message:=common.Message{}
		messageData1:=messageData
		json.Unmarshal(messageData1,&message)

		//初始化client的群频道		//同时让client注册到这个频道里 模拟初次登录
		if message.UserId!="" {
			c.SenderId=message.UserId
			//群注册
			hubs:=c.InitClientHubMapAndRegister(message.UserId)
			c.HubMap=hubs
			AllClient[c.SenderId]=c
		}

		//创群
		if message.Type==common.CreateGroupType{
			c.SendCreateGroupMessage(messageData)
		}

		if (message!=common.Message{})&&message.Type==common.SingleMessageType {

			c.SendSingleMessage(messageData)
		} else if (message!=common.Message{})&&message.Type==common.GroupMessageType {

			c.SendGroupMessage(messageData)
		}else {
			c.SendComonMessage(messageData)
		}

	}
}
func(c *ClientService)SendComonMessage(messageData[]byte)  {

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
		message.SenderId=c.SenderId
		fmt.Println("此时发送者是",c.SenderId)
		//记录到历史消息里去
		c.RecordInGroupHistory(message)

		groupMessage,_:=json.Marshal(message)
		groupId:=message.GroupId
		hub:=c.HubMap[groupId]

	//检查该client是否关注了这个群
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

	c.RecordInSingleHistory(message)
	fmt.Println("私人信息",message)


	//检查收信息的人是否在线，如果不在就88 在就发送给他！
	if AllClient[message.AccepterId]!=nil{
		AllClient[message.AccepterId].AcceptedMessages<-messageData
	}


}

//初始化client的hubMap和register到hub里去
func (c *ClientService)InitClientHubMapAndRegister(userId string)map[string]*Hub {
	groups:= repository.GetGroupsByUserId(userId)
	hubMap:=make(map[string]*Hub)
	for _,v:=range groups{
		groupId:=v.GroupId
		hubMap[groupId]= AllHub[groupId]
		AllHub[groupId].Register<-c
		//获取每个群的历史信息
		c.SendGroupHistoryToClient(groupId)
		c.SendSingleHistoryToClient(c.SenderId)
	}

	return hubMap
}


func (c *ClientService)RecordInGroupHistory(message common.GroupMessage){
	historyMessage:=model.GroupHistoryMessage{
		SenderId: message.SenderId,
		GroupId: message.GroupId,
		Content: message.Content,
		Time: time.Now(),
	}
	groupHistory.RecordInHistory(&historyMessage)

}
func (c *ClientService)RecordInSingleHistory(message common.SingleMessage){
	historyMessage:=model.SingleHistoryMessage{
		SenderId: c.SenderId,
		AccepterId: message.AccepterId,
		Content: message.Content,
		Time: time.Now(),
	}
	singleHistory.RecordInHistory(&historyMessage)

}

func (c *ClientService)SendGroupHistoryToClient(groupId string){
	groupHistory:=groupHistory.GetHistory(groupId)
	hub:=AllHub[groupId]
	if hub!=nil {
		for _,v:=range groupHistory{
			historymessage,_:=json.Marshal(v)
				hub.GroupBroadCast<-historymessage
		}
	}


}

func (c *ClientService)SendSingleHistoryToClient(userId string){
	//自己和别人的会话？acctper和sender都是自己
	singleHistory:=singleHistory.GetHistory(userId)
	for _,historymessage:=range  singleHistory{
		historymessageByte,_:=json.Marshal(historymessage)
		c.AcceptedMessages<-historymessageByte

	}
}

func (c *ClientService)SendCreateGroupMessage(messageData[]byte) {
	createGroup:=common.CreateGroupMessage{}
	json.Unmarshal(messageData,&createGroup)
	c.CreateGroup(c.SenderId,createGroup.UserIds,createGroup.GroupName)
}
//创建群聊
func (c *ClientService)CreateGroup(createrId string,groupUsers []string,groupName string,)  {
	groupId:= uuid.NewV4().String()
	tx:=common.WebsocketDB.Begin()
	//建群
	err:=repository.CreateGroup(tx,groupId,groupName,createrId)
	if err!=nil {
		fmt.Errorf(err.Error())

		tx.Rollback()
		return
	}

	//加入群记录
	err=repository.AddUsersToGroup(tx,groupUsers,groupId)
	if err!=nil {
		fmt.Errorf(err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()

	///创建群控制器
	hub:=NewHub(groupId)
	go hub.Run()
	hub.CreateGroupInit(groupId,groupUsers)
	//在总观这里注册
	AllHub[groupId]=hub

}

//todo 注销 还没写好
func (c *ClientService)LogOut(){
	AllClient[c.SenderId]=nil
	for _,hub:=range c.HubMap {
		hub.Unregister<-c
	}
}