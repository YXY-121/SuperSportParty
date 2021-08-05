package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

type SingleHistoryMessageRepository struct {

}

func (s *SingleHistoryMessageRepository)RecordInHistory(singleMessage *model.SingleHistoryMessage)  {
	common.DB.Model(&model.SingleHistoryMessage{}).Create(&singleMessage)
}

func (s *SingleHistoryMessageRepository)GetHistory(userId string) []model.SingleHistoryMessage {
	singleHistory:=make([]model.SingleHistoryMessage,0)
	common.DB.Model(&model.SingleHistoryMessage{}).Where("sender_id=? or accepter_id=?",userId,userId).Scan(&singleHistory)
	return singleHistory
}
