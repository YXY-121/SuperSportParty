package repository

import (
	"apiproject/common"
	"apiproject/model"
	"fmt"

	"github.com/sirupsen/logrus"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

//根据位置获取订单大概详情？ 不用把所有的值都返回回去
func (o *OrderRepository) GetOrdersByLocation(location string) []model.Order {
	orders := make([]model.Order, 0)
	//order:=model.Order{}

	common.SportDB.Model(&model.Order{}).Where("order_location=?", location).Scan(&orders)
	return orders
}

//获取单个订单的信息
func (o *OrderRepository) GetOrderByOrderId(orderId string) model.Order {
	order := model.Order{}
	common.SportDB.Model(&model.Order{}).Where("order_id=?", orderId).Scan(&order)
	return order

}

//获取单个订单的信息
func (o *OrderRepository) GetOrderByTypeAndLocation(orderType string, location string) []model.Order {

	order := make([]model.Order, 0)
	common.SportDB.Model(&model.Order{}).Where("sport_type=?", orderType).Where("order_location=?", location).Find(&order)
	return order

}

//创建订单
func (o *OrderRepository) CreateOrder(order model.Order) {
	err := common.SportDB.Model(&model.Order{}).Create(order).Error
	if err != nil {
		logrus.Errorln("create order fail")
		return
	}
}

func (o *OrderRepository) DelOrderPeopleNumber(orderId string, num int) {
	//先查询是否>0，不然就88
	order := model.Order{}
	//如果==1，说明满人了就拉人进群里
	fmt.Println(11)

	common.SportDB.Model(&model.Order{}).Scan(&order).Where("order_id=?", orderId)
	fmt.Println(order)

	if order.RestNumber == 0 {
		return
	}
	if order.RestNumber == 1 {
		//满了,准备拉群
	}
	common.SportDB.Model(&model.Order{}).Where("order_id=?", orderId).Update("total_number", order.RestNumber-1)

}
