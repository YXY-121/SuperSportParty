package model

type Group struct {
	GroupId      string `gorm:"group_id" json:"group_id"`
	GroupName    string `gorm:"group_name"  json:"group_name"`
	GroupCreater string `gorm:"group_creater"  json:"group_creater"`
}

func (Group) TableName() string {
	return "group"
}
