package repository

import (
	"apiproject/common"
	"apiproject/model"
)

type UserOrderRepository struct {
}

func (u *UserOrderRepository) CreateUserOrderRecord(userId, orderId string) {
	order := model.UserOrder{
		OrderId: orderId,
		UserId:  userId,
	}
	common.SportDB.Model(&model.UserOrder{}).Create(&order)
}
