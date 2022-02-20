package model

import "time"

//管理员审核
type Review struct {
	adminId          string    `gorm:"admin_id" `
	UserId           string    `gorm:"user_id" `
	OrderId          string    `gorm:"order_id" `
	IsReview         int       `gorm:"is_review" `
	ReviewContent    string    `gorm:"review_Content" `
	IsBreakTheLow    int       `gorm:"is_break_the_low" `
	ReviewCreateTime time.Time `gorm:"review_create_time" `
	ReviewUpdateTime time.Time `gorm:"review_update_time" `
}

func (Review) TableName() string {
	return "review"
}
