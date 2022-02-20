package repository

import (
	"apiproject/common"
	"apiproject/model"
)

type GroupHistoryMessageRepository struct {
}

// 插入群组消息
func RecordGroupInHistory(groupMessage *model.GroupHistoryMessage) {
	common.WebsocketDB.Model(&model.GroupHistoryMessage{}).Create(&groupMessage)
}

// 根据群id查询5条最新的历史消息
func GetGroupHistory(groupId string) []model.GroupHistoryMessage {
	groupHistory := make([]model.GroupHistoryMessage, 0)
	common.WebsocketDB.Model(&model.GroupHistoryMessage{}).Where("group_id=?", groupId).Order("time desc").Limit(5).Scan(&groupHistory)
	return groupHistory
}
