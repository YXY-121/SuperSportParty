package service

import (
	"apiproject/redPaper/model"
	"apiproject/redPaper/repository"
)


type EnvelopeItemService struct {
	repository *repository.EnvelopeItemRepository
}
//还剩多少个红包、有多少人抢了的详情

func (e *EnvelopeItemService)GetItems(envelopeId string)[]model.EnvelopeItem  {
	return 	e.GetItems(envelopeId)
}

