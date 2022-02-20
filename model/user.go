package model

type User struct {
	UserId          string `gorm:"user_id" json:"user_id"`
	UserName        string `gorm:"user_name;primaryKey" json:"user_name"`
	Phone           string `gorm:"phone;primaryKey" json:"phone"`
	Password        string `gorm:"password" json:"password"`
	City            string `gorm:"city" json:"city"`
	ChatHead        string `gorm:"chat_head" json:"chat_head"`
	SelfDescription string `gorm:"self_description" json:"self_description"`
}

func (User) TableName() string {
	return "user"
}
