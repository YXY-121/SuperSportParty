package service

import (
	"apiproject/common"
	"apiproject/model"
	"errors"
	"regexp"

	"gorm.io/gorm"
)

//检查是否已经存在该用户
func IsExistUser(phone string) bool {
	err := common.SportDB.Model(&model.User{}).Where("phone=?", phone).First(&model.User{}).Error
	// Check if returns RecordNotFound error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true

}

// 识别手机号码
func IsCorrectPhoneType(mobile string) bool {
	regular := "/^(((13[0-9]{1})|186|159|(15[0-9]{1}))/d{8})$/;"

	result, _ := regexp.MatchString(regular, mobile)
	if !result {
		return false
	}
	return true
}
