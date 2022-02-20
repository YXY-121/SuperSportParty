package service

import "apiproject/repository"

type UserOrderService struct {
	UserOrderRepository *repository.UserOrderRepository
}

func NewUserOrder() *UserOrderService {
	return &UserOrderService{
		&repository.UserOrderRepository{},
	}
}
