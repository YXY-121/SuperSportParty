package repository

import (
	"apiproject/common"
	"apiproject/redPaper/model"
)

type EnvelopeItemRepository struct {


}
func (e *EnvelopeItemRepository)GetItems(envelopeId string)[]model.EnvelopeItem  {
	itemArray:=make([]model.EnvelopeItem,0)
	common.HongbaoDB.Model(&model.EnvelopeItem{}).Where("envelope_id=?",envelopeId).Scan(&itemArray)
	return itemArray
}
