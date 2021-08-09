package repository

import "github.com/shopspring/decimal"

type AccountRepository struct {

}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}
func (a *AccountRepository)CheckMoney(userId string)  {

}
func (a *AccountRepository)PayForEnvelope(userId string,money decimal.Decimal )  {

}