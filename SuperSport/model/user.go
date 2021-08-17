package model
type User struct {
	UserId string `gorm:"user_id"`
}
func  (User)TableName() string {
	return "user"
}

