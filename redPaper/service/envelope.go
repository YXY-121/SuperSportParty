package service

import (
	"apiproject/common"
	"apiproject/redPaper/repository"
	"github.com/shopspring/decimal"
)


type EnvelopeService struct {
	repository *repository.EnvelopeRepository
}
type EnvelopeItem struct {
	userId string
	userName string
	Money decimal.Decimal
}


var accountService = NewAccountService()
func (e *EnvelopeService)CreateEnvelope(userId string,totalMoney decimal.Decimal,number int,envelopeType string)  {
	//检查这个用户的账户是否有钱
	//这里是一个事务
	tx:=common.HongbaoDB.Begin()

	if accountService.CheckIfHaveEnoughMoney(userId){
		//创建红包订单

		e.repository.CreateEnvelope(userId,totalMoney,number,envelopeType)

		//个人扣钱
		accountService.PayForEnvelope(userId,totalMoney)

		//将红包放入redis里

		tx.Commit()
	}else{
		tx.Rollback()
		//返回，你没钱！
	}

}

func (e *EnvelopeService)ReceiveEnvelope()  {
	//查看当前当但

}
//剩下的个数
func (e *EnvelopeService)GetRestNumber(envelopeId string)  {

}

