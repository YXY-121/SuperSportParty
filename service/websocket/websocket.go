package service

import (
	"apiproject/common"
	"apiproject/model"
	"apiproject/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var groupHistory = repository.GroupHistoryMessageRepository{}
var singleHistory = repository.SingleHistoryMessageRepository{}

type ClientService struct {
	Url              url.URL
	Conn             *websocket.Conn
	HubMap           map[string]*Hub
	AcceptedMessages chan []byte
	UserId           string
}

func (c *ClientService) ReturnAccectped() {
	defer c.Conn.Close()
	for {
		select {
		//通过通道发现收到了消息，就会返回给自己的客户端页面 writeMessage
		case message, ok := <-c.AcceptedMessages:

			if !ok {
				logrus.Printf("接受信息失败，err 是%v\n", ok)
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)

		}
	}

}

func (c *ClientService) SendOther() {

	for {
		messageType, messageData, err := c.Conn.ReadMessage()
		//这个很重要 不然client退出的时候 server会报错
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//注销
		if messageType == websocket.CloseMessage {
			fmt.Println("注销")
			c.LogOut()
		}

		message := common.Message{}
		messageData1 := messageData
		json.Unmarshal(messageData1, &message)

		//
		if (message == common.Message{}) {
			c.SendComonMessage(messageData)
		}

		if message.Type == common.SingleMessageType {

			c.SendSingleMessage(messageData)
		}

		if message.Type == common.GroupMessageType {

			c.SendGroupMessage(messageData)
		}

	}
}
func (c *ClientService) SendComonMessage(messageData []byte) {

	for _, hub := range c.HubMap {
		a := "group" + hub.HubId
		a = a + string(messageData[:])

		fmt.Println(string(messageData[:]))
		bytemessage := []byte(a)
		hub.GroupBroadCast <- bytemessage
	}
}

func (c *ClientService) SendGroupMessage(messageData []byte) {

	message := common.GroupMessage{}
	json.Unmarshal(messageData, &message)
	fmt.Println(message)

	message.SenderId = c.UserId
	fmt.Println("此时发送者是", c.UserId)
	//记录到历史消息里去
	c.RecordInGroupHistory(message)

	groupMessage, _ := json.Marshal(message)
	groupId := message.GroupId
	hub := c.HubMap[groupId]

	//检查该client是否关注了这个群
	if hub != nil {
		fmt.Println("已发送给hub处理")
		hub.GroupBroadCast <- groupMessage
	} else {
		fmt.Println("查不到这个群")
	}
	return

}

func (c *ClientService) SendSingleMessage(messageData []byte) {

	message := common.SingleMessage{}

	json.Unmarshal(messageData, &message)

	//写到历史记录上
	c.RecordInSingleHistory(message)
	fmt.Println("私人信息", message)

	//检查收信息的人是否在线，如果不在就88 在就发送给他！
	if AllClient[message.AccepterId] != nil {
		AllClient[message.AccepterId].AcceptedMessages <- messageData
	}

}

//初始化该用户 所对应的群hub
func (c *ClientService) InitClientHubMapAndRegister(userId string) map[string]*Hub {
	groups := repository.GetGroupsByUserId(userId)
	if groups == nil {
		return nil
	}
	hubMap := make(map[string]*Hub)

	for _, v := range groups {
		groupId := v.GroupId
		hubMap[groupId] = AllHub[groupId]
		AllHub[groupId].Register <- c
		//获取每个群的历史信息
		c.SendGroupHistoryToClient(groupId)
	}

	return hubMap
}

//RecordInGroupHistory 记录历史群聊记录
func (c *ClientService) RecordInGroupHistory(message common.GroupMessage) {
	historyMessage := model.GroupHistoryMessage{
		SenderId:   message.SenderId,
		GroupId:    message.GroupId,
		Content:    message.Content,
		CreateTime: time.Now().Unix(),
	}
	repository.RecordGroupInHistory(&historyMessage)

}

//RecordInSingleHistory 记录单人聊天记录
func (c *ClientService) RecordInSingleHistory(message common.SingleMessage) {
	historyMessage := model.SingleHistoryMessage{
		SenderId:   c.UserId,
		AccepterId: message.AccepterId,
		Content:    message.Content,
		CreateTime: time.Now().Unix(),
	}
	repository.RecordSingleInHistory(&historyMessage)

}

func (c *ClientService) SendGroupHistoryToClient(groupId string) {
	groupHistory, err := repository.GetGroupHistory(groupId)
	if err != nil {

	}
	hub := AllHub[groupId]
	if hub != nil {
		for _, v := range groupHistory {
			historymessage, _ := json.Marshal(v)
			hub.GroupBroadCast <- historymessage
		}
	}

}

func (c *ClientService) SendSingleHistoryToClient(userOneId, userTwoId string) {
	//自己和别人的会话？acctper和sender都是自己
	singleHistory, err := repository.GetSingleHistory(userOneId, userTwoId)
	if err != nil {

	}
	for _, historymessage := range singleHistory {
		historymessageByte, _ := json.Marshal(historymessage)
		c.AcceptedMessages <- historymessageByte

	}
}

//创建群聊
func CreateGroup(createrId string, groupUsers []string, groupName string) error {
	groupId := uuid.NewV4().String()
	tx := common.WebsocketDB.Begin()
	//建群
	err := repository.CreateGroup(tx, groupId, groupName, createrId)
	if err != nil {
		logrus.Errorln(err.Error())

		tx.Rollback()
		return err
	}

	//加入群记录
	err = repository.AddUsersToGroup(groupUsers, groupId)
	if err != nil {
		logrus.Errorln(err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()

	///创建群控制器
	hub := NewHub(groupId)
	go hub.Run()
	hub.CreateGroupInit(groupId, groupUsers)
	//在总观这里注册
	AllHub[groupId] = hub
	return nil

}

//todo 注销 还没写好
func (c *ClientService) LogOut() {
	AllClient[c.UserId] = nil
	for _, hub := range c.HubMap {
		hub.Unregister <- c
	}
}

//心跳检测
