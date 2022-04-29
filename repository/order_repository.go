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
func GetOrderByTypeAndLocation(orderType string, latitude, longtitue *float64) []model.Order {

	selectSql := ""
	selectCondition := make([]interface{}, 0)
	var sqlValue = make(map[string]interface{})
	sqlValue["sport_type"] = orderType

	keys := make([]string, 0, len(sqlValue))
	for k := range sqlValue {
		keys = append(keys, k)
	}

	for index, val := range keys {
		if index == 0 {
			selectSql += "where " + val + "=?"

		} else {
			selectSql += " and where " + val + "=?"
		}
		selectCondition = append(selectCondition, sqlValue[val].(string))

	}

	fmt.Println(selectCondition)
	fmt.Println(len(selectCondition))
	order := make([]model.Order, 0)
	common.SportDB.Model(&model.Order{}).Select(selectSql, selectCondition...).Scan(&order)
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
func DeleteOrder(orderId string) error {
	order := model.Order{}
	return common.SportDB.Delete(&order, "order_id", orderId).Error

}
func ListOrders() ([]model.Order, error) {
	orders := make([]model.Order, 0)

	err := common.SportDB.Model(&model.Order{}).Scan(&orders).Error

	return orders, err
}

func ListOrdersByLongitudeAndLatitud(longtitue, latitude float64) ([]model.Order, error) {
	orders := make([]model.Order, 0)

	err := common.SportDB.Model(&model.Order{}).
		Where("longtitue=?", longtitue).
		Where("latitude=?", latitude).
		Scan(&orders).Error
	return orders, err
}
func ListUnreviewOrder() ([]model.Order, error) {
	orders := make([]model.Order, 0)
	err := common.SportDB.Model(&model.Order{}).
		Where("is_review=?", 0).
		Scan(&orders).Error
	return orders, err
}
func UpdateOrderStatus(orderId string, isReview int) error {
	err := common.SportDB.Model(&model.Order{}).
		Where("order_id=?", orderId).
		Update("is_review", isReview).
		Error

	return err
}
func ListAllUserByOrder(orderId string) ([]model.User, error) {
	var users = make([]model.User, 0)
	err := common.SportDB.Model(model.User{}).Where("order_id=?", orderId).Scan(&users).Error
	return users, err
}

// func ListOrderByUserId(userId string) {
// 	var orders = make([]model.Order, 0)
// 	err := common.SportDB.Model(model.User{}).Where("user_id=?", userId).Scan(&orders).Error

//}
