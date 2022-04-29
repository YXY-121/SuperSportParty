package repository

import (
	"apiproject/common"
	"apiproject/model"
	"fmt"

	"gorm.io/gorm"
)

// 获取全部的组信息
func GetAllGroups() []model.Group {
	group := make([]model.Group, 0)
	common.WebsocketDB.Model(model.Group{}).Scan(&group)
	return group
}

//创建群聊
func CreateGroup(tx *gorm.DB, groupId string, groupName string, groupMaster string) error {
	group := model.Group{
		groupId,
		groupName,
		groupMaster,
	}
	err := tx.Model(&model.Group{}).Create(&group).Error
	if err != nil {
		fmt.Errorf(err.Error())
		tx.Rollback()
		return err
	}
	return nil
}
