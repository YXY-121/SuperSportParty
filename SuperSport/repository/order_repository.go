package repository

import (
	"apiproject/SuperSport/model"
	"apiproject/common"
	"gorm.io/gorm"
)

type OrderRepository struct {


}
func (o *OrderRepository)GetOrderByLocation()  {

}
func (o *OrderRepository)AddOrder(order model.Order)  {
		common.SportDB.Model(&model.Order{}).Save(order)

}
func (o *OrderRepository)DelOrderPeopleNumber(orderId string,num int)  {
	//先查询是否>0，不然就88
	order:=model.Order{}
	//如果==1，说明满人了就拉人进群里
	common.SportDB.Model(&model.Order{}).Scan(&order).Where("order_id=?",orderId)
	if (order!=model.Order{}){
		if order.RestNumber==0 {
			return
		}
		if order.RestNumber==1{
			//满了,准备拉群
		}
		common.SportDB.Model(&model.Order{}).Update("people_number",gorm.Expr("people_number-?",num)).Where("order_id=?",orderId)

	}

}