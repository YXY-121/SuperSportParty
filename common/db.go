package common

import (
	"apiproject/config"
	"apiproject/model"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var WebsocketDB *gorm.DB
var HongbaoDB *gorm.DB
var SportDB *gorm.DB
var RedisDB redis.Conn

func InitDB() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, config.App.Database.User, config.App.Database.Pwd, config.App.Database.Host, config.App.Database.Name)
	//	dsn1 := "root:123@tcp(127.0.0.1:3306)/chat?charset=utf8&parseTime=True&loc=Local"
	//	dsn2 := "root:123@tcp(127.0.0.1:3306)/hongbao?charset=utf8&parseTime=True&loc=Local"
	//	dsn3 := "root:123@tcp(127.0.0.1:3306)/supersport?charset=utf8&parseTime=True&loc=Local"
	var err error
	SportDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	SportDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&model.Order{}, &model.Sport{}, &model.User{},
		&model.UserOrder{}, &model.UserAppraisal{},
		&model.Group{}, &model.GroupHistoryMessage{}, &model.SingleHistoryMessage{}, &model.UserGroup{})
	WebsocketDB = SportDB
	//	WebsocketDB, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	//	HongbaoDB, err = gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err != nil {
		logrus.Println(err)
	}

}
func initRedis() {
	var err error
	RedisDB, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

}
