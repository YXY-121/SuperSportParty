package model

import "time"

type Order struct {
	CreaterId string `gorm:"creater_id"`
	OrderLocation string `gorm:"order_location"`
	SportType string `gorm:"sport_type"`
	Description string `gorm:"description"`
	PeopleNumber int `gorm:"people_number"`
	CreateTime time.Time `gorm:"create_time"`
	EndTime time.Time `gorm:"end_time"`

}
func  (Order)TableName() string {
	return "order"
}