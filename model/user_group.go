package model

type UserGroup struct {
	UserId  string `gorm:"user_id;primary_key"`
	GroupId string `gorm:"group_id;primary_key"`
}

func (UserGroup) TableName() string {
	return "user_group"
}
