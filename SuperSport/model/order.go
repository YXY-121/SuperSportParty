package model

import "time"

type Order struct {
	OrderId string `gorm:"order_id"`
	CreaterId string `gorm:"creater_id"`
	SportType string `gorm:"sport_type"`
	Description string `gorm:"description"`
	TotalNumber int `gorm:"total_number"`
	RestNumber int `gorm:"rest_number"`
	OrderLocation string`gorm:"order_location"`
	Longitude string `gorm:"longitude"`
	Latitude string `gorm:"latitude"`
	CreateTime time.Time `gorm:"create_time"`
	EndTime time.Time `gorm:"end_time"`

}
func  (Order)TableName() string {
	return "order"
}