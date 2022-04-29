package repository

import (
	"apiproject/common"
	"apiproject/model"
	"fmt"
)

//获取某人的所有好友
func ListFriendByUserId(userID int) ([]model.UserHeadAndName, error) {
	friendsId := make([]int, 0)
	//获取所有的朋友id
	err := common.SportDB.Model(&model.UserFriend{}).Raw("? UNION ?",
		common.SportDB.Model(&model.UserFriend{}).Select("user_id1").Where("user_id2=?", userID),
		common.SportDB.Model(&model.UserFriend{}).Select("user_id2").Where("user_id1=?", userID),
	).Scan(friendsId).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	friends := make([]model.UserHeadAndName, 0)
	//头像和id
	err = common.SportDB.Model(&model.User{}).Where("id IN (?)", friendsId).
		Scan(friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil

}

//新增好友
func AddFriend(ID, friendID int) error {
	if ID > friendID {
		tmpID := friendID
		friendID = ID
		ID = tmpID
	}
	friend := model.UserFriend{}
	friend.UserId1 = ID
	friend.UserId2 = friendID
	return common.SportDB.Model(&model.UserFriend{}).Create(&friend).Error
}

//删除好友
func DeleteFriend(ID, friendID int) error {
	if ID > friendID {
		tmpID := friendID
		friendID = ID
		ID = tmpID
	}
	friend := model.UserFriend{}
	friend.UserId1 = ID
	friend.UserId2 = friendID
	return common.SportDB.Model(&model.UserFriend{}).Delete(&friend).Error

}

func HaveBeenFriend(ID, friendID int) bool {
	userFriend := model.UserFriend{}
	if ID > friendID {
		tmpID := friendID
		friendID = ID
		ID = tmpID
	}
	err := common.WebsocketDB.Model(&model.UserFriend{}).Where("user_id1=?", ID).Where("user_id2=?", friendID).
		Scan(&userFriend).Error

	if (err != nil || userFriend == model.UserFriend{}) {
		return false
	}
	return true
}
