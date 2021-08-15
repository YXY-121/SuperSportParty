package common

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var WebsocketDB *gorm.DB
var HongbaoDB *gorm.DB
var SportDB *gorm.DB
var RedisDB redis.Conn
func init()  {
	initDB()
	initRedis()
}
func initDB (){
	dsn1 := "root:123@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	dsn2 := "root:123@tcp(127.0.0.1:3306)/hongbao?charset=utf8mb4&parseTime=True&loc=Local"
	dsn3:="root:123@tcp(127.0.0.1:3306)/supersport?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	SportDB,err=gorm.Open(mysql.Open(dsn3), &gorm.Config{})
	WebsocketDB, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	HongbaoDB,err=gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err!=nil {
		log.Println(err)
	}
}
func initRedis()  {
	var err error
	RedisDB, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")



}
