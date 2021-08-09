package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Account struct {
	AccountId string 	 `gorm:"account_id"`
	AccountType string `gorm:"account_type"`
	UserId string  `gorm:"user_id"`
	AccountBalance decimal.Decimal `gorm:"account_balance"`
	Status int `gorm:"status"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}


func  (Account)TableName() string {
	return "account"
}


