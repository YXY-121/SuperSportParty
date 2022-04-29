package model

type User struct {
	ID              int    `gorm:"primaryKey"`
	UserId          string `gorm:"user_id" json:"user_id`
	UserName        string `gorm:"user_name;primaryKey" json:"user_name"`
	Phone           string `gorm:"phone;" json:"phone"`
	Password        string `gorm:"password" json:"password"`
	City            string `gorm:"city" json:"city"`
	ChatHead        string `gorm:"chat_head" json:"chat_head"`
	SelfDescription string `gorm:"self_description" json:"self_description"`
	IsLock          int    `gorm:"is_lock",json:"is_lock"`
}

func (User) TableName() string {
	return "user"
}

type UserHeadAndName struct {
	UserName string `gorm:"user_name;primaryKey" json:"user_name"`
	ChatHead string `gorm:"chat_head" json:"chat_head"`
	UserId   string `gorm:"user_id" json:"user_id"`
}
