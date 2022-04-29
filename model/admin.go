package model

type Admin struct {
	AdminId string `json:"admin_id" gorm:"admin_id"`
	Pwd     string `json:"pwd" gorm:"pwd"`
}

func (Admin) TableName() string {
	return "admin"
}
