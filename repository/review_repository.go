package repository

import (
	"apiproject/common"
	"apiproject/model"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func ListReviewByAdmin(adminId string, startSince, startUntil int64, isReview int) ([]model.Review, error) {
	req := make([]model.Review, 0)

	if err := common.SportDB.Model(model.Review{}).Where("admin_id=?", adminId).Scan(&req).Error; err != nil {
		logrus.Errorln("ListReviewByAdmin err is :", err)
		return nil, err
	}
	return req, nil

}

//封号处理
//如果发现又要新增一条reviwe的时候，选出所有违法的帖子。如果违法的帖子已经>=2条了，还没到 就封号吧
func CloseAccountByReview(userId string) error {
	var count int64
	//查询个数
	if err := common.SportDB.Model(model.Review{}).Group(userId).Count(&count).Error; err != nil {
		logrus.Errorln("CloseAccountByReview err is :", err)
		return nil
	}
	if count >= 2 {
		if err := common.SportDB.Model(model.User{}).Update("is_lock", 1).Where("user_id=?", userId).Error; err != nil {
			return err
		}
	}
	return nil

}
func ListIllegalReview() []model.Review {
	req := make([]model.Review, 0)
	if err := common.SportDB.Model(model.Review{}).Where("is_break_the_low=?", 1).Scan(&req).Error; err != nil {
		logrus.Errorln("ListReviewByAdmin err is :", err)
		return nil
	}
	return req
}

func ListReviewByUserIdAndTime(userId string, createSince, createUntil int64, isReview int) []model.Review {
	req := make([]model.Review, 0)
	err := common.SportDB.Model(model.Review{}).
		Where("user_id=?", userId).
		Where("create_time >=?", createSince).
		Where("create_time <=?", createUntil).
		Scan(&req).Error
	if err != nil {
		logrus.Errorln("ListReviewByAdmin err is :", err)
		return nil
	}
	return req
}

func CreateReview(req model.Review) error {
	//判断是否已经存在该待审核的id
	if err := common.SportDB.Model(model.Review{}).Create(&req).Error; err != nil {
		return err
	}
	return nil

}
func UpdateReviewRecord(req model.Review) error {

	if err := common.SportDB.Model(model.Review{}).Create(&req).Error; err != nil {
		return err
	}
	return nil
}
func ListReviewByOrderId(orderID string) ([]model.Review, error) {
	req := make([]model.Review, 0)
	if err := common.SportDB.Model(model.Review{}).Where("order_id=?", orderID).Order("update_time desc").Scan(&req).Error; err != nil {
		return nil, err
	}
	return req, nil
}
func GetCreateTime(orderId, userId string) (int64, error) {
	var req model.Review
	if err := common.SportDB.Model(model.Review{}).
		Where("order_id=?", orderId).
		Where("user_id=?", userId).Find(&req).Error; err != nil {
		return 0, err
	}
	return req.CreateTime, nil
}
func CreateNullReview(adminId, orderId, userId string) error {
	//获取userid orderid
	var req model.Review
	req.ReviewId = uuid.NewV4().String()
	req.OrderId = orderId
	req.UserId = userId
	req.IsReview = 0
	return common.SportDB.Model(model.Review{}).Create(&req).Error
}
