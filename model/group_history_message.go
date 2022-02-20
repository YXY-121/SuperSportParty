package model

import "time"

type GroupHistoryMessage struct {
	SenderId string    `gorm:"sender_id"`
	Content  string    `gorm:"content"`
	GroupId  string    `gorm:"group_id"`
	Time     time.Time `gorm:"time"`
}

func (GroupHistoryMessage) TableName() string {
	return "group_history_message"
}
