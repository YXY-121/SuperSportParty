package model

type Group struct {
	GroupId string 	 `gorm:grou_id`
}


func  (Group)TableName() string {
	return "group"
}


