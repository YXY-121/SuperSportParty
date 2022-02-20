package client

import (
	"apiproject/websocket/repository"
	"fmt"
)

var AllHub = make(map[string]*Hub)              //当前所有的群
var AllClient = make(map[string]*ClientService) //在线的user对应的AllClient.如果不在线就为nil

type Hub struct {
	HubId          string
	Register       chan *ClientService //上线就注册到这个频道里
	Unregister     chan *ClientService
	Clients        map[*ClientService]bool //当前群在线人员
	GroupBroadCast chan []byte             //信息发送渠道

}

//初始化全部的group
func InitAllGroup() {
	groups := repository.GetAllGroups()
	for _, group := range groups {
		hub := NewHub(group.GroupId)
		AllHub[group.GroupId] = hub
		go hub.Run()
	}
	fmt.Println("当前的群总数", len(AllHub))
}

func NewHub(hubId string) *Hub {
	return &Hub{
		HubId:          hubId,
		Register:       make(chan *ClientService), //通知要注册的频道s
		Unregister:     make(chan *ClientService),
		Clients:        make(map[*ClientService]bool),
		GroupBroadCast: make(chan []byte),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			//在线
			h.Clients[client] = true

		case message := <-h.GroupBroadCast:
			for client := range h.Clients {
				select {
				case client.AcceptedMessages <- message:
				default:
					close(client.AcceptedMessages)
					delete(h.Clients, client)
				}
			}
		case client := <-h.Unregister:
			fmt.Println(client.SenderId, "跑路啦！")
			h.Clients[client] = false
		}

	}
}
func (h *Hub) CreateGroupInit(groupId string, groupUsers []string) {
	//将在线的给拉群里 不在线的话，到时候上线就会自动进群
	for _, user := range groupUsers {
		client := AllClient[user]
		if client != nil {
			client.HubMap[groupId] = h
			h.Register <- client
			fmt.Println(client.SenderId, "的群数量：", len(client.HubMap))

		}

	}

}
