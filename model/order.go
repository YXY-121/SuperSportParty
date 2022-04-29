package model

import "apiproject/pkg/page"

type Order struct {
	OrderId     string  `gorm:"order_id" json:"order_id"`
	CreaterId   string  `gorm:"creater_id json:"creater_id"`
	SportType   string  `gorm:"sport_type" json:"sport_type"`
	Description string  `gorm:"description" json:"description"`
	TotalNumber int     `gorm:"total_number" json:"total_number"`
	RestNumber  int     `gorm:"rest_number" json:"rest_number"`
	Longitude   float64 `gorm:"longitude" json:"longitude"`
	Latitude    float64 `gorm:"latitude" json:"latitude"`
	CreateTime  int64   `gorm:"create_time" json:"create_time"`
	EndTime     int64   `gorm:"end_time" json:"end_time"`
	LockTime    string  `gorm:"lock_time" `
	ReviewId    string  `gorm:"review_id" `
	Location    string  `gorm:"location" json:"location" `
	IsReview    int     `gorm:"is_review" json:"is_review"`
	PicUrls     string  `gorm:"pic_urls" json:"pic_urls" `
}
type OrderReq struct {
	CreaterId   string   `json:"creater_id" form:"creater_id"`
	Description string   ` json:"description" form:"description"`
	TotalNumber *int     ` json:"total_number" form:"total_number"`
	Longitude   *float64 `json:"longitude" form:"longitude"`
	Latitude    *float64 `json:"latitude" form:"latitude"`
	page.Page
	OrderType string `json:"sport_type" form:"sport_type"`
	PicUrls   string `gorm:"pic_urls" json:"pic_urls" `
}

func (Order) TableName() string {
	return "order"
}
