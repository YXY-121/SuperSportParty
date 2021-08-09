package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
	"fmt"
	"gorm.io/gorm"
)

func GetAllGroups() []model.Group {
	group:=make([]model.Group,0)

	common.WebsocketDB.Model(model.Group{}).Scan(&group)
	return  group
}
func GetAllUsersByGroupId(groupId string)  {

	common.WebsocketDB.Model(model.Group{}).Where("group_id=?",groupId)

}
func CreateGroup(tx *gorm.DB,groupId string, groupName string, groupMaster string)error {
	group:=model.Group{
		groupId,
		groupName,
		groupMaster,
	}
	err:=tx.Model(&model.Group{}).Create(&group).Error
	if err!=nil {
		fmt.Errorf(err.Error())
		tx.Rollback()
		return err
	}
	return  nil
}
