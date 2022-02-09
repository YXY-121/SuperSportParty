package model

import "time"

type Order struct {
	OrderId       string    `gorm:"order_id",json:"order_id"`
	CreaterId     string    `gorm:"creater_id,json:"creater_id"`
	SportType     string    `gorm:"sport_type",json:"sport_type"`
	Description   string    `gorm:"description",json:"description"`
	TotalNumber   int       `gorm:"total_number",json:"total_number"`
	RestNumber    int       `gorm:"rest_number",json:"rest_number"`
	OrderLocation string    `gorm:"order_location",json:"order_location"`
	Longitude     string    `gorm:"longitude",json:"longitude"`
	Latitude      string    `gorm:"latitude",json:"latitude"`
	CreateTime    time.Time `gorm:"create_time",json:"create_time"`
	EndTime       time.Time `gorm:"end_time"`
	StringEndTime string    `json:"endTime"`
}

func (Order) TableName() string {
	return "order"
}
