package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

type GroupHistoryMessageRepository struct {

}


func (s *GroupHistoryMessageRepository)RecordInHistory(groupMessage *model.GroupHistoryMessage)  {

	common.WebsocketDB.Model(&model.GroupHistoryMessage{}).Create(&groupMessage)
}
func (s *GroupHistoryMessageRepository)GetHistory(groupId string) []model.GroupHistoryMessage {
	groupHistory:=make([]model.GroupHistoryMessage,0)
	common.WebsocketDB.Model(&model.GroupHistoryMessage{}).Where("group_id=?",groupId).Order("time desc").Limit(5).Scan(&groupHistory)
	return groupHistory
}
