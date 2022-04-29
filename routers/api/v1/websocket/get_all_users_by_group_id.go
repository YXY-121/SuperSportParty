package v1

import (
	"apiproject/pkg"
	"apiproject/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GroupIdReq struct {
	GroupId string `json:"group_id"`
}

// 根据groupid获取组里的成员信息
func GetAllUsersByGroupId(c *gin.Context) {
	var req GroupIdReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	users, err := repository.GetAllUsersByGroupId(req.GroupId)
	if err != nil {
		pkg.Response(c, pkg.ErrorListFail, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	pkg.Response(c, pkg.Success, users)
}
