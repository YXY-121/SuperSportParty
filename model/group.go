package model

type Group struct {
	GroupId string 	 `gorm:"group_id"`
	GroupName string `gorm:"group_name"`
	GroupMaster string `gorm:"group_master"`
}


func  (Group)TableName() string {
	return "group"
}



