package v1

import (
	"apiproject/pkg"
	"apiproject/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//根据用户id获取其相应的所有群信息
func GetGroupsByUserId(c *gin.Context) {
	var userId UserId
	err := c.ShouldBind(&userId)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	groups := repository.GetGroupsByUserId(userId.UserId)

	pkg.Response(c, pkg.Success, groups)
}
