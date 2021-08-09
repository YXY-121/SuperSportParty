package model

import (
	"github.com/shopspring/decimal"
	"time"

)


type Envelope struct {
	EnvelopeId string 	 `gorm:"envelope_id"`
	EnvelopeType string  `gorm:"envelope_type"`
	SenderId string  `gorm:"sender_id"`
	TotalMoney decimal.Decimal `gorm:"total_money"`
	RestMoney decimal.Decimal `gorm:"rest_money"`
	Number int `gorm:"number"`
	ExpiredTime time.Time `gorm:"expired_time"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`

}


func  (Envelope)TableName() string {
	return "envelope"
}

