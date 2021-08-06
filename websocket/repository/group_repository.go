package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

func GetAllGroups() []model.Group {
	group:=make([]model.Group,0)

	common.DB.Model(model.Group{}).Scan(&group)
	return  group
}
func GetAllUsersByGroupId(groupId string)  {

	common.DB.Model(model.Group{}).Where("group_id=?",groupId)

}
func CreateGroup(groupId string,groupName string,groupMaster string){
	group:=model.Group{
		groupId,
		groupName,
		groupMaster,
	}

	common.DB.Model(&model.Group{}).Create(&group)
}
