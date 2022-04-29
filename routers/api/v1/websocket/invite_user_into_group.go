package v1

import (
	"apiproject/pkg"
	"apiproject/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GroupId struct {
	GroupID string `json:"group_id"`
}

type InviteUserReq struct {
	GroupID string `json:"group_id"`

	UserArray []string `json:"user_array"`
}

//拉某人进群
func InviteUserInGroup(c *gin.Context) {

	var req InviteUserReq
	err := c.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	err = repository.AddUsersToGroup(req.UserArray, req.GroupID)
	if err != nil {
		logrus.Errorf("邀请进群失败")
		pkg.Response(c, pkg.ErrorUpdateFail, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	logrus.Errorf("邀请进群成功")
	pkg.Response(c, pkg.Success, nil)

}
