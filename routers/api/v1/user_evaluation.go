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

type EvaluationReq struct {
	page.Page
	UserId
}

//创建评价
func CreateEvaluation(c *gin.Context) {
	var req model.Evaluation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	req.EvaluationId = uuid.NewV4().String()
	req.CreateTime = time.Now().Unix()
	err = repository.CreateEvaluation(req)
	if err != nil {
		logrus.Errorln("CreateEvaluation err :", err)
		pkg.Response(c, pkg.ErrorCreateFail, nil)
		return
	}
	pkg.Response(c, pkg.Success, "创建成功")

}

//查看该用户的所有评价
func GetEvaluationByUserId(c *gin.Context) {
	var req EvaluationReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	evaluations, err := repository.ListEvaluation(req.UserId.UserId)
	//通过userid和orderid来看？可以
	if err != nil {
		logrus.Println("ListEvaluation fail,err:", err)
		pkg.Response(c, pkg.ErrorListEvaluationFail, pkg.MsgMap[pkg.ErrorListEvaluationFail])
		return
	}
	pkg.Response(c, pkg.Success, evaluations)
}

//查看某个人的综合评价，获取全部star 平均
//根据大家的评价，来计算出这个人的总体信誉评分。
func CalculateUserLevel() {

}
func ListAllUserByOrder(c *gin.Context) {
	var req OrderID
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	users, err := repository.ListAllUserByOrder(req.OrderId)
	if err != nil {
		logrus.Println("ListAllUserByOrder fail,err:", err)
		pkg.Response(c, pkg.ErrorListOrderFail, pkg.MsgMap[pkg.ErrorListOrderFail])
		return
	}
	pkg.Response(c, pkg.Success, users)

}
