package repository

import (
	"apiproject/common"
	"apiproject/model"

	"github.com/sirupsen/logrus"
)

type UserGroupRepository struct {
}

// 根据用户id获取其所属的群聊
func GetGroupsByUserId(userId string) []model.UserGroup {
	group := make([]model.UserGroup, 0)
	err := common.WebsocketDB.Model(model.UserGroup{}).Where("user_id=?", userId).Scan(&group).Error
	if err != nil {
		logrus.Errorln("查找群聊失败")
		return nil
	}
	return group
}

// 将用户加入群聊 前段做一个限定，只有不在群的才能显示可以拉群
func AddUsersToGroup(users []string, groupId string) error {

	userGroup := make([]model.UserGroup, 0)
	for _, v := range users {
		temp := model.UserGroup{}
		temp.GroupId = groupId
		temp.UserId = v
		userGroup = append(userGroup, temp)
	}
	return common.WebsocketDB.Model(&model.UserGroup{}).Create(&userGroup).Error

}

// 根据groupid获取组里的成员信息
func GetAllUsersByGroupId(groupId string) ([]model.UserHeadAndName, error) {
	users := make([]model.UserHeadAndName, 0)
	userIds := make([]string, 0)
	err := common.WebsocketDB.Model(&model.UserGroup{}).Select("user_id").Where("group_id=?", groupId).Scan(&userIds).Error
	if err != nil {
		return nil, err
	}
	err = common.WebsocketDB.Model(&model.User{}).Where("user_id IN (?)", userIds).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
