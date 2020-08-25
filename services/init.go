package services

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/recardoz/cusim/models"
)

//todo 數據庫配置信息
const (
	DB_HOST = "127.0.0.1"
	DB_PORT = "3306"
	DB_USER = "root"
	//DB_PASS    = "Recardo-0126"
	//DB_PASS	= "root"
	DB_PASS    = "Gcbody-0865"
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
	//db.LogMode(DB_DEBUG)

	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Users{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Chats{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Emails{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.IpUsers{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Apps{})

	logs.Informational("连接数据库成功")
}
