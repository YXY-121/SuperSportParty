package v1

import (
	service "apiproject/service/websocket"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

type UserId struct {
	UserId string `json:"user_id"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
} // use default options
func checkOrigin(r *http.Request) bool {
	return true
}

func ServerWs(c *gin.Context) {

	//升级成websocket协议
	//每次都创建一个client

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	userId := c.Request.FormValue("user_id")
	client := &service.ClientService{
		Conn:             conn,
		AcceptedMessages: make(chan []byte, 256),
		UserId:           userId,
	}
	//初始化client的群频道		//同时让client注册到这个频道里 模拟初次登录
	//群注册
	fmt.Println("userid:", userId)
	fmt.Println(service.AllHub)
	hubs := client.InitClientHubMapAndRegister(userId)
	client.HubMap = hubs
	service.AllClient[client.UserId] = client

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	if err != nil {
		log.Println("read:", err)
		return
	}

	go client.ReturnAccectped()
	go client.SendOther()

}

type CreateGroupReq struct {
	CreaterId  string   `json:"creater_id"`
	GroupUsers []string `json:"group_users"`
	GroupName  string   `json:"group_name"`
}

//删除聊天记录这个以后再加
