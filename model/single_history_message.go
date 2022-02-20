package model

import "time"

type SingleHistoryMessage struct {
	SenderId       string    `gorm:"sender_id"`
	Content        string    `gorm:"content"`
	AccepterId     string    `gorm:"accepter_id"`
	IsAccepterRead int       `gorm:"is_accepter_read"`
	Time           time.Time `gorm:"time"`
}

func (SingleHistoryMessage) TableName() string {
	return "single_history_message"
}
