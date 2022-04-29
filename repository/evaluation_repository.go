package repository

import (
	"apiproject/common"
	"apiproject/model"
)

func CreateEvaluation(evaluation model.Evaluation) error {
	return common.WebsocketDB.Model(&model.Evaluation{}).Create(&evaluation).Error

}
func ListEvaluation(userId string) ([]model.Evaluation, error) {
	evaluations := make([]model.Evaluation, 0)
	err := common.WebsocketDB.Model(&model.Evaluation{}).Where("user_id=?", userId).Scan(evaluations).Error
	return evaluations, err
}
