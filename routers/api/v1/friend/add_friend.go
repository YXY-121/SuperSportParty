package v1

import (
	"apiproject/model"
	"apiproject/pkg"
	"apiproject/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Friend struct {
	IsFriend int
	User     model.User
}

//增加好友
func AddFriend(c *gin.Context) {
	userID, err := strconv.Atoi(c.Request.FormValue("ID"))
	friendID, err := strconv.Atoi(c.Request.FormValue("friend_ID"))

	if userID == friendID || err != nil {
		//参数错误
		pkg.Response(c, pkg.ErrorParams, nil)
		return
	}
	//判断是否为好友，如果是的话直接展示详细信息以及能够发消息的窗口
	if repository.HaveBeenFriend(userID, friendID) {
		//返回该好友的详细信息
		friend := Friend{}
		friend.User, err = repository.GetUserByID(friendID)
		if err != nil {
			pkg.Response(c, pkg.ErrorCreateFail, nil)
			return
		}
		pkg.Response(c, pkg.Success, friend)
		return
	}

	//
	err = repository.AddFriend(userID, friendID)
	if err != nil {
		pkg.Response(c, pkg.ErrorCreateFail, nil)
		return

	}

	pkg.Response(c, pkg.Success, "添加好友成功")

}

//删除好友
func DeleteFriend(c *gin.Context) {
	userID, err := strconv.Atoi(c.Request.FormValue("ID"))
	friendID, err := strconv.Atoi(c.Request.FormValue("friend_ID"))

	if userID == friendID || err != nil {
		//参数错误
		pkg.Response(c, pkg.ErrorParams, nil)
		return
	}
	err = repository.DeleteFriend(userID, friendID)
	if err != nil {
		pkg.Response(c, pkg.ErrorListFail, nil)
		return
	}
	pkg.Response(c, pkg.Success, nil)

}

//查看全部好友
func ListFriend(c *gin.Context) {
	userID, err := strconv.Atoi(c.Request.FormValue("ID"))

	if err != nil {
		//参数错误
		pkg.Response(c, pkg.ErrorParams, nil)
		return
	}
	friends, err := repository.ListFriendByUserId(userID)
	if err != nil {
		pkg.Response(c, pkg.ErrorListFail, nil)
		return
	}
	pkg.Response(c, pkg.Success, friends)

}

//查看某个好友
func GetFriendById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Request.FormValue("ID"))
	friendID, err := strconv.Atoi(c.Request.FormValue("friend_ID"))

	if err != nil {
		//参数错误
		pkg.Response(c, pkg.ErrorParams, nil)
		return
	}
	//判断是否为好友
	if !repository.HaveBeenFriend(userID, friendID) {
		pkg.Response(c, pkg.ErrorListFail, nil)
		return
	}
	//是的话就获取所有信息
	friends, err := repository.GetUserByID(friendID)
	if err != nil {
		pkg.Response(c, pkg.ErrorListFail, nil)
		return
	}
	pkg.Response(c, pkg.Success, friends)

}
