package reponsitory

import (
	"apiproject/common"
	"apiproject/websocket/model"
)

func GetAllGroups() []model.Group {
	group:=make([]model.Group,0)

	common.DB.Model(model.Group{}).Scan(&group)
	return  group
}
