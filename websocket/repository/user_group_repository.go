package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
	"fmt"
)

type UserGroupRepository struct {

}

func GetGroupsByUserId(userId string) []model.UserGroup {
	group:=make([]model.UserGroup,0)

	common.DB.Model(model.UserGroup{}).Where("user_id=?",userId).Scan(&group)
	return  group
}
func AddUsersToGroup(groupUsers []string,groupId string){
	group:=make([]model.UserGroup,0)
	for _,v:=range groupUsers{
		temp:=model.UserGroup{}
		temp.GroupId=groupId
		temp.UserId=v
		group=append(group, temp)
	}
	common.DB.Model(&model.UserGroup{}).Create(&group)
	fmt.Println("创群",groupId)
}
