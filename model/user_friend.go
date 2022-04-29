package model

type UserFriend struct {
	UserId1 int `gorm:"user_id1;primary_key"`
	UserId2 int `gorm:"user_id2;primary_key"`
}

func (UserFriend) TableName() string {
	return "user_friend"
}
