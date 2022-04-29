package v1

import (
	"apiproject/pkg"
	service "apiproject/service/websocket"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//创建群聊
func CreateGroup(c *gin.Context) {

	var req CreateGroupReq
	err := c.ShouldBind(&req)
	logrus.Println(req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	err = service.CreateGroup(req.CreaterId, req.GroupUsers, req.GroupName)
	if err != nil {
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
	}
	pkg.Response(c, pkg.Success, nil)

}
