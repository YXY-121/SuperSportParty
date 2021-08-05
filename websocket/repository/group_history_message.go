package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

type GroupHistoryMessageRepository struct {

}


func (s *GroupHistoryMessageRepository)RecordInHistory(groupMessage *model.GroupHistoryMessage)  {

	common.DB.Model(&model.GroupHistoryMessage{}).Create(&groupMessage)
}
func (s *GroupHistoryMessageRepository)GetHistory(groupId string) []model.GroupHistoryMessage {
	groupHistory:=make([]model.GroupHistoryMessage,0)
	common.DB.Model(&model.GroupHistoryMessage{}).Where("group_id=?",groupId).Scan(&groupHistory)
	return groupHistory
}
