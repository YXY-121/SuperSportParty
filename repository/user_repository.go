package repository

import (
	"apiproject/common"
	"apiproject/model"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateUser(user model.User) error {
	err := common.SportDB.Model(&model.User{}).Create(&user).Error
	if err != nil {
		logrus.Errorln("create user fail")
		return err
	}
	return nil
}

func GetUser(userId string) (model.User, error) {
	user := model.User{}
	err := common.SportDB.Model(&model.User{}).Where("user_id=?", userId).Find(&user).Error
	logrus.Println(err)
	logrus.Println(errors.Is(err, gorm.ErrRecordNotFound))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("get user fail")
		return user, err
	}
	return user, nil

}
func GetUserByID(userID int) (model.User, error) {
	user := model.User{}
	err := common.SportDB.Model(&model.User{}).Where("ID=?", userID).Find(&user).Error
	logrus.Println(err)
	logrus.Println(errors.Is(err, gorm.ErrRecordNotFound))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("get user fail")
		return user, err
	}
	return user, nil

}

func UpdateUser(user model.User) error {
	err := common.SportDB.Model(&model.User{}).Where("user_id=?", user.UserId).Save(user).Error
	if err != nil {
		logrus.Errorln("update user fail")
		return err
	}
	return nil
}
func DeleteUser(userId string) error {
	return common.SportDB.Model(&model.User{}).Where("user_id=?", userId).Delete(&model.User{}).Error
}
func UpdateUserLock(userId string, lock int) error {
	err := common.SportDB.Model(&model.User{}).Where("user_id=?", userId).Update("is_lock", lock).Error
	if err != nil {
		logrus.Errorln("update user lock status fail")
		return err
	}
	return nil
}
