package service

import "apiproject/SuperSport/repository"

type UserOrderService struct {
	UserOrderRepository *repository.UserOrderRepository
}

func NewUserOrder() *UserOrderService {
	return &UserOrderService{
		&repository.UserOrderRepository{},
	}
}