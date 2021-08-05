package repository

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

func GetGroupsByUserId(userId string) []model.UserGroup {
	group:=make([]model.UserGroup,0)

	common.DB.Model(model.UserGroup{}).Where("user_id=?",userId).Scan(&group)
	return  group
}