package model

type User struct {
	UserId          string `gorm:"user_id"`
	UserName        string `gorm:"user_name"`
	City            string `gorm:"city"`
	ChatHead        string `gorm:"chat_head"`
	SelfDescription string `gorm:"self_description"`
	IsLock          int    `gorm:"is_lock"`
}

func (User) TableName() string {
	return "user"
}
