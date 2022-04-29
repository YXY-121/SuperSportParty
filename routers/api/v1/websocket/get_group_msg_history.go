package v1

import (
	"apiproject/pkg"
	"apiproject/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//获取群聊历史聊天记录
func GetGroupMsgHistory(c *gin.Context) {
	var req GroupIdReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	history, err := repository.GetGroupHistory(req.GroupId)

	if err != nil {
		logrus.Errorf("GetGroupHistory fail")
		pkg.Response(c, pkg.ErrorListFail, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	pkg.Response(c, pkg.Success, history)

}
