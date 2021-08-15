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
		common.RedisDB.Do("set","1","2")

		tx.Commit()
	}else{
		tx.Rollback()
		//返回，你没钱！
	}

}

//func (e *EnvelopeService)ReceiveEnvelope(envelope model.Envelope)  decimal.Decimal{
//	//获取当前的剩的钱，然后计算出金额
//	if envelope.RestNumber==1{
//		return envelope.RestMoney
//	}
//
//	if envelope.Type=="1" {
//		return envelope.TotalMoney
//	}
//
//
//}
//func randomMoney(restAmount decimal.Decimal, restNumber int) decimal.Decimal{
//	decimalNumber:=decimal.NewFromInt(int64(restNumber))
//	decimalRand:=restAmount.Div(decimalNumber).Mul(decimal.NewFromInt(int64(2)).Sub(decimal.NewFromInt(int64(1)))).Add(decimal.NewFromInt(int64(1)))
//	amount := rand.Intn(int(decimalRand))
//
//
//	//	restAmount.Sub(amount)
//	v:=decimal.NewFromInt(int64(amount))
//	fmt.Println(v)
//	fmt.Println(amount)
//
//	return v
//}


//剩下的个数
func (e *EnvelopeService)GetRestNumber(envelopeId string)  {

}

