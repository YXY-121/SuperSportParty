package v1

import (
	"apiproject/model"
	"apiproject/pkg"
	"apiproject/pkg/page"
	"apiproject/repository"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type AdminID struct {
	AdminId string `json:"admin_id"`
}
type ReviewReq struct {
	page.Page
	AdminID
	UserId
	IsReview   int   `json:"is_review"`
	StartSince int64 `json:"start_since"`
	StartUntil int64 `json:"start_until"`
}
type ReviewOrderReq struct {
	OrderID
	page.Page
}
type MyReponse struct {
	page.Page
	Data interface{} `json:"data"`
}

//查看管理员的审核记录
func GetReviewRecordByAdmin(c *gin.Context) {
	//
	var req ReviewReq
	err := c.ShouldBindJSON(&req)
	if err != nil || req.PageSize <= 0 || req.PageNo <= 0 {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	reviews, err := repository.ListReviewByAdmin(req.AdminId, req.StartSince, req.StartUntil, req.IsReview)
	if err != nil {
		logrus.Errorln("GetReviewRecordByAdmin err :", err)
		pkg.Response(c, pkg.ErrorListFail, nil)
		return
	}
	flag, start, end := page.PageTool(len(reviews), req.PageSize, req.PageNo)
	if !flag {
		//分页失败
		pkg.Response(c, pkg.ErrorPageFail, pkg.MsgMap[pkg.ErrorPageFail])
		return
	}

	req.Page.Total = len(reviews)
	pkg.Response(c, pkg.Success, MyReponse{req.Page, reviews[start:end]})

}

//查看用户的审核情况 超过4次发恶劣的帖子就被封号处理
func ListReviewByUserIdAndTime(c *gin.Context) {
	var req ReviewReq
	err := c.ShouldBindJSON(&req)
	if err != nil || req.PageSize <= 0 || req.PageNo <= 0 {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	reviews := repository.ListReviewByUserIdAndTime(req.UserId.UserId, req.StartSince, req.StartUntil, req.IsReview)
	flag, start, end := page.PageTool(len(reviews), req.PageSize, req.PageNo)
	if !flag {
		//分页失败
		pkg.Response(c, pkg.ErrorPageFail, pkg.MsgMap[pkg.ErrorPageFail])
		return
	}

	req.Page.Total = len(reviews)
	pkg.Response(c, pkg.Success, MyReponse{req.Page, reviews[start:end]})
}

//创建审核记录
func CreateReviewRecord(c *gin.Context) {
	var req model.Review
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	req.CreateTime = time.Now().Unix()
	req.UpdateTime = time.Now().Unix()
	req.ReviewId = uuid.NewV4().String()
	if err := repository.CreateReview(req); err != nil {
		logrus.Errorln("CreateReviewRecord err :", err)
		pkg.Response(c, pkg.ErrorCreateFail, nil)
		return
	}
	pkg.Response(c, pkg.Success, "创建成功")

}

//修改审核记录
func UpdateReviewRecord(c *gin.Context) {
	//这里只修改update time不修改create time
	var req model.Review
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	req.ReviewId = uuid.NewV4().String()
	req.UpdateTime = time.Now().Unix()
	createTime, err := repository.GetCreateTime(req.OrderId, req.UserId)
	if err != nil {
		req.CreateTime = req.UpdateTime
	} else {
		req.CreateTime = createTime
	}
	req.IsReview = 1
	if err := repository.UpdateReviewRecord(req); err != nil {
		logrus.Errorln("UpdateReviewRecord err :", err)
		pkg.Response(c, pkg.ErrorUpdateFail, nil)
		return
	}
	pkg.Response(c, pkg.Success, "修改成功")

}

//通过同一个orderid，查询所有审核记录。因为可能会有出现修改审核的记录。
func ListReviewByOrderId(c *gin.Context) {
	//这里只修改update time不修改create time
	var req ReviewOrderReq
	err := c.ShouldBindJSON(&req)
	if err != nil || req.PageSize <= 0 || req.PageNo <= 0 {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	reviews, err := repository.ListReviewByOrderId(req.OrderId)
	if err != nil {
		logrus.Errorln("UpdateReviewRecord err :", err)
		pkg.Response(c, pkg.ErrorUpdateFail, nil)
		return
	}
	flag, start, end := page.PageTool(len(reviews), req.PageSize, req.PageNo)
	if !flag {
		//分页失败
		pkg.Response(c, pkg.ErrorPageFail, pkg.MsgMap[pkg.ErrorPageFail])
		return

	}
	req.Page.Total = len(reviews)
	pkg.Response(c, pkg.Success, MyReponse{req.Page, reviews[start:end]})

}
func ListUnreviewOrder(c *gin.Context) {
	var req page.Page
	err := c.ShouldBindJSON(&req)
	if err != nil || req.PageSize <= 0 || req.PageNo <= 0 {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	orders, err := repository.ListUnreviewOrder()
	if err != nil {

	}

	flag, start, end := page.PageTool(len(orders), req.PageSize, req.PageNo)
	if !flag {
		//分页失败
		pkg.Response(c, pkg.ErrorPageFail, pkg.MsgMap[pkg.ErrorPageFail])
		return

	}
	req.Total = len(orders)
	pkg.Response(c, pkg.Success, MyReponse{req, orders[start:end]})
}
