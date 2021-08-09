package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
	"fmt"
	"gorm.io/gorm"
)

type UserGroupRepository struct {

}

func GetGroupsByUserId(userId string) []model.UserGroup {
	group:=make([]model.UserGroup,0)

	common.WebsocketDB.Model(model.UserGroup{}).Where("user_id=?",userId).Scan(&group)
	return  group
}
func AddUsersToGroup(tx *gorm.DB,groupUsers []string,groupId string)error{
	group:=make([]model.UserGroup,0)
	for _,v:=range groupUsers{
		temp:=model.UserGroup{}
		temp.GroupId=groupId
		temp.UserId=v
		group=append(group, temp)
	}
	fmt.Println("创群",groupId)

	return tx.Model(&model.Group{}).Create(&group).Error

}
