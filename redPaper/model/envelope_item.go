package model

import "github.com/shopspring/decimal"

type EnvelopeItem struct {
	EnvelopeId string `gorm:"envelope_id"`
	EnvelopeItemId string `gorm:"envelope_item_id"`
	Money decimal.Decimal `gorm:"money"`
	AccepterId string `gorm:"accepter_id"`
}