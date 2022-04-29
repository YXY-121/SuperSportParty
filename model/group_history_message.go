package model

type GroupHistoryMessage struct {
	SenderId   string `gorm:"sender_id"`
	Content    string `gorm:"content"`
	GroupId    string `gorm:"group_id"`
	CreateTime int64  `gorm:"create_time"`
}

func (GroupHistoryMessage) TableName() string {
	return "group_history_message"
}
