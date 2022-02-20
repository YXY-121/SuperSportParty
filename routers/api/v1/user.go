package v1

import (
	"apiproject/config"
	"apiproject/middleware"
	"apiproject/model"
	"apiproject/pkg"
	"apiproject/repository"
	"apiproject/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type UserId struct {
	UserId string `json:"user_id"`
}

func Login(c *gin.Context) {
	g := pkg.Gin{C: c}
	var req UserId
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	token := middleware.CreateJwt(req.UserId)
	c.Writer.Header().Set(config.App.Server.JwtHeader, token)

	g.Response(http.StatusOK, pkg.Success, nil)

}

//创建用户
func CreateUser(c *gin.Context) {
	g := pkg.Gin{C: c}

	var user model.User
	err := c.ShouldBind(&user)
	//检查手机格式
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	//检查手机格式
	if service.IsCorrectPhoneType(user.Phone) {
		logrus.Errorf("手机格式错误!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	//检查手机是否重复
	if service.IsExistUser(user.Phone) {
		g.Response(http.StatusInternalServerError, pkg.ErrorUserHaveExist, nil)
		return
	}

	user.UserId = uuid.NewV1().String()
	err = repository.CreateUser(user)
	if err != nil {
		g.Response(http.StatusInternalServerError, pkg.Error, nil)
		return
	}
	g.Response(http.StatusOK, pkg.Success, user)

}

//修改用户信息
func UpdateUser(c *gin.Context) {
	g := pkg.Gin{C: c}

	var user model.User
	err := c.ShouldBind(&user)
	//检查各个参数是否正确，比如手机长度
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	if service.IsCorrectPhoneType(user.Phone) {
		logrus.Errorf("手机格式错误!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	err = repository.UpdateUser(user)
	if err != nil {
		g.Response(http.StatusInternalServerError, pkg.Error, nil)
		return
	}
	g.Response(http.StatusOK, pkg.Success, user)
}

//锁定非法用户
func UpdateLockUser(c *gin.Context) {
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
	err = repository.UpdateUserLock(userId.UserId)
	if err != nil {
		g.Response(http.StatusInternalServerError, pkg.Error, nil)
		return
	}
	g.Response(http.StatusOK, pkg.Success, nil)
}

//获取用户信息
func GetUserInfo(c *gin.Context) {
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
	user, flag := repository.GetUser(userId.UserId)
	if !flag {
		g.Response(http.StatusInternalServerError, pkg.Error, nil)
		return
	}
	g.Response(http.StatusOK, pkg.Success, user)
}

//删除用户
func DeleteUser(c *gin.Context) {
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
	err = repository.DeleteUser(userId.UserId)
	if err != nil {
		g.Response(http.StatusInternalServerError, pkg.Error, nil)
		return
	}
	g.Response(http.StatusOK, pkg.Success, nil)
}
