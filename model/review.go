package model

//管理员审核
type Review struct {
	ReviewId      string `gorm:"review_id" json:"review_id"`
	AdminId       string `gorm:"admin_id" json:"admin_id"`
	UserId        string `gorm:"user_id" json:"user_id"`
	OrderId       string `gorm:"order_id" json:"order_id"`
	ReviewContent string `gorm:"review_content" json:"review_content"`
	IsBreakTheLow int    `gorm:"is_break_the_low" json:"is_break_the_low"`
	CreateTime    int64  `gorm:"create_time" json:"create_time"`
	UpdateTime    int64  `gorm:"update_time" json:"update_time"`
	IsReview      int    `gorm:"is_review" json:"is_review"`
}

func (Review) TableName() string {
	return "review"
}
