package model

type UserGroup struct {
	UserId string 	`gorm:"user_id"`
	GroupId string `gorm:"group_id"`

}

func  (UserGroup)TableName() string {
	return "user_group"
}
