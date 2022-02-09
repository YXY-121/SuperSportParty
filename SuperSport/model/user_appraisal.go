package model

import "time"

type UserAppraisal struct {
	Reviewer string    `gorm:"reviewer"`
	UserId   string    `gorm:"user_id"`
	Content  string    `gorm:"content"`
	Level    string    `gorm:"level"`
	Time     time.Time `gorm:"time"`
}

func (UserAppraisal) TableName() string {
	return "user_appraisal"
}
