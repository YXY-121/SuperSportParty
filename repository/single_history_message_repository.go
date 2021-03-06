package repository

import (
	"apiproject/common"
	"apiproject/model"
)

type SingleHistoryMessageRepository struct {
}

// 插入个人聊天信息
func RecordSingleInHistory(singleMessage *model.SingleHistoryMessage) {
	common.WebsocketDB.Model(&model.SingleHistoryMessage{}).Create(&singleMessage)
}

// TODO 获取单人聊天记录改用http完成
// 获取两个人之间的消息记录
func GetSingleHistory(userOneId string, userTwoId string) ([]model.SingleHistoryMessage, error) {
	singleHistory := make([]model.SingleHistoryMessage, 0)

	err := common.WebsocketDB.Model(&model.SingleHistoryMessage{}).Where(
		common.WebsocketDB.Model(&model.SingleHistoryMessage{}).Where("sender_id=? or accepter_id=?", userOneId, userTwoId).
			Or(
				common.WebsocketDB.Model(&model.SingleHistoryMessage{}).Where("sender_id=? or accepter_id=?", userTwoId, userOneId),
			)).Order("time desc").
		Limit(5).
		Scan(&singleHistory).Error

	return singleHistory, err
}
