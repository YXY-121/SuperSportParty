package model

type UserGroup struct {
	UserId string
	GroupId string

}

func  (UserGroup)TableName() string {
	return "user_group"
}
