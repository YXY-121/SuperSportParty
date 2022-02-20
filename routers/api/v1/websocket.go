package v1

import (
	"apiproject/pkg"
	client "apiproject/websocket/client/service"
	Repository "apiproject/websocket/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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

	client := &client.ClientService{
		Conn:             conn,
		AcceptedMessages: make(chan []byte, 256),
	}

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

//根据用户id获取其相应的所有群信息
func GetGroupsByUserId(c *gin.Context) {
	g := pkg.Gin{C: c}

	var userId UserId
	err := c.ShouldBind(&userId)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	groups := Repository.GetGroupsByUserId(userId.UserId)

	g.Response(http.StatusOK, pkg.Success, groups)
}

type GroupIdReq struct {
	GroupId string `json:"group_id"`
}

// 根据groupid获取组里的成员信息
func GetAllUsersByGroupId(c *gin.Context) {
	g := pkg.Gin{C: c}
	var req GroupIdReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	users, err := Repository.GetAllUsersByGroupId(req.GroupId)
	if err != nil {
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
	}

	g.Response(http.StatusOK, pkg.Success, users)
}

type CreateGroupReq struct {
	CreaterId  string   `json:"creater_id"`
	GroupUsers []string `json:"group_users"`
	GroupName  string   `json:"group_name"`
}

//创建群聊
func CreateGroup(c *gin.Context) {
	g := pkg.Gin{C: c}
	var req CreateGroupReq
	err := c.ShouldBind(&req)
	logrus.Println(req)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	err = client.CreateGroup(req.CreaterId, req.GroupUsers, req.GroupName)
	if err != nil {
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
	}
	g.Response(http.StatusOK, pkg.Success, nil)

}

//获取群聊历史聊天记录
func GetGroupMsgHistory(c *gin.Context) {
	g := pkg.Gin{C: c}
	var req GroupIdReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	Repository.GetGroupHistory(req.GroupId)
	g.Response(http.StatusOK, pkg.Success, nil)

}

//获取用户之间的历史聊天记录
func GetUserMsgHistory(c *gin.Context) {

	g := pkg.Gin{C: c}
	var req GroupIdReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	Repository.GetGroupHistory(req.GroupId)
	g.Response(http.StatusOK, pkg.Success, nil)
}

//拉某人进群
func InviteUserInGroup(c *gin.Context) {

}

//删除聊天记录这个以后再加
