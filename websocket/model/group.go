package model

type Group struct {
	GroupId string 	 `gorm:"group_id"`
}


func  (Group)TableName() string {
	return "group"
}


