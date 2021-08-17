package model


type UserOrder struct {
	OrderId string `gorm:"order_id"`
	UserId string `gorm:"user_id"`
}
func  (UserOrder)TableName() string {
	return "user_order"
}
