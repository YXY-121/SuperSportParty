package model
type Sport struct {
	SportType string `gorm:"sport_type"`
	SportName string `gorm:"sport_name"`
}
func  (Sport)TableName() string {
	return "sport"
}

