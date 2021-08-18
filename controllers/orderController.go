package controllers

import (
	"apiproject/SuperSport/model"
	"apiproject/SuperSport/service"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type OrderController struct {
	beego.Controller
}
var  orderService=service.NewOrderService()
//get
func (o*OrderController)GetOrderById()  {
	orderId := o.GetString("orderId")
	fmt.Println(orderId)
	order:=orderService.GetOrderById(orderId)
	o.Data["json"] = order
	o.ServeJSON()

}
//get
func (o*OrderController)GetOrdersByLoacation()  {
	location := o.GetString("location")
	fmt.Println(location)
	order:=orderService.GetOrderByLocation(location)
	o.Data["json"] = order
	o.ServeJSON()


}
//post
func (o*OrderController)CreateOrder(){
	order:=model.Order{}
	//这是json
	json.Unmarshal(o.Ctx.Input.RequestBody,&order)
	fmt.Println(order)
	//createrId string,orderLocation string,sportType string,description string,peopleNumber int,longitude,latitude string ,endTime time.Time)
	orderService.CreateOrder(order.CreaterId,order.OrderLocation,order.SportType,order.Description,order.TotalNumber,order.Longitude,order.Latitude,order.EndTime)
	o.Data["json"] = order
	o.ServeJSON()
}

type GrabOrder struct {
	OrderId string
	UserId string
}
//post
func (o*OrderController)GrabOrder(){
	grabOrder:=GrabOrder{}
	fmt.Println(o.Ctx.Input.RequestBody)
	//这是json
	json.Unmarshal(o.Ctx.Input.RequestBody,&grabOrder)
	fmt.Println("g",grabOrder)
	if (grabOrder==GrabOrder{}) {
		o.Data["json"] = nil
		o.ServeJSON()

	}
	//createrId string,orderLocation string,sportType string,description string,peopleNumber int,longitude,latitude string ,endTime time.Time)
	orderService.GrabOrder(grabOrder.UserId,grabOrder.OrderId)
	o.Data["json"] = nil
	o.ServeJSON()
}

