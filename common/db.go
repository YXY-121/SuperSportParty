package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
func init()  {
		dsn := "root:123@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil {
		log.Println(err)
	}

}