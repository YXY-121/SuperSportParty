package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

type SingleHistoryMessageRepository struct {

}

func (s *SingleHistoryMessageRepository)RecordInHistory(singleMessage *model.SingleHistoryMessage)  {
	common.WebsocketDB.Model(&model.SingleHistoryMessage{}).Create(&singleMessage)
}

func (s *SingleHistoryMessageRepository)GetHistory(userId string) []model.SingleHistoryMessage {
	singleHistory:=make([]model.SingleHistoryMessage,0)
	common.WebsocketDB.Model(&model.SingleHistoryMessage{}).Where("sender_id=? or accepter_id=?",userId,userId).Order("time desc").Limit(5).Scan(&singleHistory)
	return singleHistory
}
