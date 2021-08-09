package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var WebsocketDB *gorm.DB
var HongbaoDB *gorm.DB
func init()  {
		dsn1 := "root:123@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	dsn2 := "root:123@tcp(127.0.0.1:3306)/hongbao?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
	WebsocketDB, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	HongbaoDB,err=gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err!=nil {
		log.Println(err)
	}

}