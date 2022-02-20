package repository

import (
	"apiproject/common"
	"apiproject/model"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateUser(user model.User) error {
	err := common.SportDB.Model(&model.User{}).Create(user).Error
	if err != nil {
		logrus.Errorln("create user fail")
		return err
	}
	return nil
}

func GetUser(userId string) (model.User, bool) {
	user := model.User{}
	err := common.SportDB.Model(&model.User{}).Where("user_id=?", userId).Find(&user).Error
	logrus.Println(err)
	logrus.Println(errors.Is(err, gorm.ErrRecordNotFound))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("get user fail")
		return user, false
	}
	return user, true

}

func UpdateUser(user model.User) error {
	err := common.SportDB.Model(&model.User{}).Where("user_id=?", user.UserId).Updates(user).Error
	if err != nil {
		logrus.Errorln("update user fail")
		return err
	}
	return nil
}
func DeleteUser(userId string) error {
	return common.SportDB.Model(&model.User{}).Where("user_id=?", userId).Delete(&model.User{}).Error
}
func UpdateUserLock(userId string) error {
	err := common.SportDB.Model(&model.User{}).Update("is_lock", userId).Error
	if err != nil {
		logrus.Errorln("update user lock status fail")
		return err
	}
	return nil
}
