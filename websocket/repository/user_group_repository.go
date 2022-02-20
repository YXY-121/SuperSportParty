package repository

import (
	"apiproject/common"
	"apiproject/model"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

//TODO AddUsersToGroup 改成http？
// 将用户加入群聊
func AddUsersToGroup(tx *gorm.DB, groupUsers []string, groupId string) error {
	group := make([]model.UserGroup, 0)
	for _, v := range groupUsers {
		temp := model.UserGroup{}
		temp.GroupId = groupId
		temp.UserId = v
		group = append(group, temp)
	}
	fmt.Println("创群", groupId)

	return tx.Model(&model.UserGroup{}).Create(&group).Error

}

// 根据groupid获取组里的成员信息
func GetAllUsersByGroupId(groupId string) ([]model.User, error) {
	users := make([]model.User, 0)
	err := common.WebsocketDB.Table("(?) as u", common.WebsocketDB.Model(&model.UserGroup{}).Select("user_id")).
		Where("group_id = ?", groupId).Scan(users).Error
	// SELECT * FROM (SELECT `name`,`age` FROM `users`) as u WHERE `age` = 18
	if err != nil {
		return nil, err
	}
	return users, nil
}
