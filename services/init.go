package services

import (
	"CustomIM/models"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//todo 數據庫配置信息
const (
	DB_HOST    = "127.0.0.1"
	DB_PORT    = "3306"
	DB_USER    = "root"
	DB_PASS    = "root"
	DB_NAME    = "cim"
	DB_CHARSET = "utf8"
	DB_DEBUG   = true
)

//todo 初始化数据库连接
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", DB_USER+":"+DB_PASS+"@("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset"+DB_CHARSET)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(DB_DEBUG)

	db.AutoMigrate(&models.Users{},&models.Chats{},&models.Emails{},&models.IpUsers{},&models.Users{},&models.Apps{})

	logs.Informational("连接数据库成功")
}
