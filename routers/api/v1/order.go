package v1

import (
	"apiproject/model"
	"apiproject/pkg"
	"apiproject/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var orderService = service.NewOrderService()

type OrderReq struct {
	OrderType string `json:"type"`
	Location  string `json:"location"`
}
type OrderIdReq struct {
	OrderId string `json:"order_id"`
}

func GetOrderById(c *gin.Context) {
	// 反序列化 请求方http body
	g := pkg.Gin{C: c}

	var orderId OrderIdReq
	err := c.ShouldBind(&orderId)
	fmt.Println("req", orderId)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return

	}
	order := orderService.GetOrderById(orderId.OrderId)
	g.Response(http.StatusOK, pkg.Success, order)
	return

}
func GetOrderByTypeAndLocation(c *gin.Context) {
	// 反序列化 请求方http body
	g := pkg.Gin{C: c}

	var req OrderReq
	err := c.ShouldBind(&req)
	fmt.Println("req", req)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return

	}
	order := orderService.GetOrderByTypeAndLocation(req.OrderType, req.Location)
	g.Response(http.StatusOK, pkg.Success, order)
	return

}

//根据地区以及喜欢的类型进行定时推送
func GetOrdersByLoacation(c *gin.Context) {
	//fmt.Println(location)
	//order := orderService.GetOrderByLocation(location)

}

func CreateOrder(c *gin.Context) {
	//这是json
	g := pkg.Gin{C: c}
	var order model.Order
	err := c.ShouldBind(&order)
	logrus.Println(order)

	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})

		return
	}
	order.CreateTime = time.Now()
	order.OrderId = uuid.NewV1().String()
	order.EndTime, err = time.Parse("2006-01-02 15:04:05", order.StringEndTime)
	//createrId string,orderLocation string,sportType string,description string,peopleNumber int,longitude,latitude string ,endTime time.Time)

	orderService.CreateOrder(order)

	g.Response(http.StatusOK, pkg.Success, nil)
	return

}

type GrabOrder struct {
	OrderId string
	UserId  string
}

//抢单
func GrabOrder1(c *gin.Context) {
	// grabOrder := GrabOrder{}
	// fmt.Println(o.Ctx.Input.RequestBody)
	// //这是json
	// json.Unmarshal(o.Ctx.Input.RequestBody, &grabOrder)
	// fmt.Println("g", grabOrder)
	// if (grabOrder == GrabOrder{}) {
	// 	o.Data["json"] = nil
	// 	o.ServeJSON()

	// }
	// //createrId string,orderLocation string,sportType string,description string,peopleNumber int,longitude,latitude string ,endTime time.Time)
	// orderService.GrabOrder(grabOrder.UserId, grabOrder.OrderId)
	// o.Data["json"] = nil
	// o.ServeJSON()
}

//删除订单
func DeleteOrder(c *gin.Context) {

}

//修改订单
func UpdateOrder(c *gin.Context) {
	//这是json
	g := pkg.Gin{C: c}
	var order model.Order
	err := c.ShouldBind(&order)
	logrus.Println(order)
	if err != nil {
		logrus.Errorf("json bind error!")
		g.Response(http.StatusBadRequest, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	order.CreateTime = time.Now()
	order.OrderId = uuid.NewV1().String()
	order.EndTime, err = time.Parse("2006-01-02 15:04:05", order.StringEndTime)
	//createrId string,orderLocation string,sportType string,description string,peopleNumber int,longitude,latitude string ,endTime time.Time)

	orderService.CreateOrder(order)

	g.Response(http.StatusOK, pkg.Success, nil)
	return
}
