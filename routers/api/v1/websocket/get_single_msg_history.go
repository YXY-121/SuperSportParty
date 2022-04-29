package v1

import (
	"apiproject/pkg"
	"apiproject/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//获取用户之间的历史聊天记录
func GetUserMsgHistory(c *gin.Context) {
	user1 := c.Request.FormValue("user_id1")
	user2 := c.Request.FormValue("user_id2")
	if user1 == "" || user2 == "" {
		logrus.Errorf("param error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%s", "缺少参数"),
		})
		return
	}
	history, err := repository.GetSingleHistory(user1, user2)
	if err != nil {
		logrus.Errorf("GetSingleHistory fail!")
		pkg.Response(c, pkg.ErrorListFail, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	pkg.Response(c, pkg.Success, history)
}
