package repository

import (
	"apiproject/common"
	"apiproject/redPaper/model"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
)

type EnvelopeRepository struct {

}

func NewEnvelopeRepository() *EnvelopeRepository {
	return &EnvelopeRepository{}
}

func (e *EnvelopeRepository)CreateEnvelope(userId string,totalMoney decimal.Decimal,number int,envelopeType string)  {
	envelope:=model.Envelope{
		EnvelopeId: uuid.NewV4().String(),
		EnvelopeType: envelopeType,
		SenderId: userId,
		TotalMoney: totalMoney,
		Number: number,
		RestMoney: totalMoney,
		ExpiredTime: time.Now().Add(24*time.Hour),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	common.HongbaoDB.Model(&model.Envelope{}).Create(&envelope)

}

type NowEnvelope struct {
Items	[]model.EnvelopeItem
RestNumber int
}
func  (e *EnvelopeRepository)GetRest()  {


}
