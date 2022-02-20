// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apiproject/middleware"
	v1 "apiproject/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//路由初始化写入日志
	router := gin.New()
	//todo ?
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.LogToGin())
	router.Use(middleware.JwtToGin())
	//	api := router.Group("/supersport").Use()

	router.POST("/login", v1.Login)

	router.POST("/get_order_by_id", v1.GetOrderById)
	router.POST("/get_order_by_type_and_location", v1.GetOrderByTypeAndLocation)
	router.POST("/get_orders_by_location", v1.GetOrdersByLoacation)
	router.POST("/grab_order", v1.GrabOrder1)
	router.POST("/create_order", v1.CreateOrder)

	router.POST("/get_user_info", v1.GetUserInfo)
	router.POST("/update_user", v1.UpdateUser)
	router.POST("/create_user", v1.CreateUser)
	router.POST("/update_lock_user", v1.UpdateLockUser)
	router.POST("/delete_user", v1.DeleteUser)

	router.GET("/ping", v1.ServerWs)
	router.POST("/create_group", v1.CreateGroup)

	router.POST("/get_groups_by_user_id", v1.GetGroupsByUserId)
	router.POST("/get_all_users_by_group_id", v1.GetAllUsersByGroupId)
	router.POST("/get_group_msg_history", v1.GetGroupMsgHistory)
	router.POST("/get_user_msg_history", v1.GetUserMsgHistory)
	router.POST("/invite_user_in_group", v1.InviteUserInGroup)

	return router
}
