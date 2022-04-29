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
	friend "apiproject/routers/api/v1/friend"
	websocket "apiproject/routers/api/v1/websocket"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//路由初始化写入日志
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.LogToGin())
	router.Use(cors.Default())

	//router.Use(middleware.JwtToGin())
	//	api := router.Group("/supersport").Use()

	router.POST("/login", v1.Login)

	//订单
	orderGroup := router.Group("order")
	orderGroup.POST("/get_order_by_id", v1.GetOrderById)
	orderGroup.POST("/get_orders_by_location_type", v1.GetOrdersByLoacation)
	orderGroup.POST("/grab_order", v1.GrabOrder1)
	orderGroup.POST("/create_order", v1.UpdateOrder)
	orderGroup.POST("/update_order", v1.UpdateOrder)
	orderGroup.POST("/list_orders", v1.ListOrders)
	orderGroup.POST("/delete_order", v1.DeleteOrder)
	//	orderGroup.POST("/upload_picture", v1.UploadPicture)

	//用户
	userGroup := router.Group("user")
	userGroup.POST("/get_user_info", v1.GetUserInfo)
	userGroup.POST("/update_user", v1.UpdateUser)
	userGroup.POST("/create_user", v1.CreateUser)
	userGroup.POST("/update_lock_user", v1.UpdateLockUser)
	userGroup.POST("/delete_user", v1.DeleteUser)

	//聊天
	chatGroup := router.Group("chat")
	chatGroup.GET("/websocket", websocket.ServerWs)
	chatGroup.POST("/create_group", websocket.CreateGroup)
	chatGroup.POST("/get_groups_by_user_id", websocket.GetGroupsByUserId)
	chatGroup.POST("/get_all_users_by_group_id", websocket.GetAllUsersByGroupId)
	chatGroup.POST("/get_group_msg_history", websocket.GetGroupMsgHistory)
	chatGroup.POST("/get_user_msg_history", websocket.GetUserMsgHistory)
	chatGroup.POST("/invite_user_in_group", websocket.InviteUserInGroup)

	//好友
	friendGroup := router.Group("friend")
	friendGroup.POST("/add_friend", friend.AddFriend)
	friendGroup.POST("/list_friends", friend.ListFriend)
	friendGroup.POST("/get_friend_by_id", friend.GetFriendById)
	friendGroup.POST("/delete_friend", friend.DeleteFriend)

	//审核
	reviewGroup := router.Group("review")
	reviewGroup.POST("/create_review_record", v1.CreateReviewRecord)
	reviewGroup.POST("/get_review_record_by_admin", v1.GetReviewRecordByAdmin)
	reviewGroup.POST("/list_review_by_user_id_and_time", v1.ListReviewByUserIdAndTime)
	reviewGroup.POST("/update_review_record", v1.UpdateReviewRecord)
	reviewGroup.POST("/list_review_by_order_id", v1.ListReviewByOrderId)
	reviewGroup.POST("/list_unreview_order", v1.ListUnreviewOrder)

	//评价
	evaluationGroup := router.Group("evaluate")
	evaluationGroup.POST("/create_evaluation", v1.CreateEvaluation)
	evaluationGroup.POST("/get_evaluation_by_userId", v1.GetEvaluationByUserId)

	return router
}
