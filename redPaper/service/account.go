package service

import (
	"apiproject/redPaper/repository"
	"github.com/shopspring/decimal"
)

type AccountService struct {
	repository *repository.AccountRepository
}

func NewAccountService()  *AccountService{
	return &AccountService{
		repository.NewAccountRepository(),
	}
}

func (e *AccountService)CheckIfHaveEnoughMoney(userId string) bool {
		e.repository.CheckMoney(userId)
		return true
}

func (e *AccountService)PayForEnvelope(userId string,money decimal.Decimal) bool {
	e.repository.CheckMoney(userId)
	return true
}
