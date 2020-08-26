package services

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/recardoz/cusim/models"
	"github.com/recardoz/cusim/utils"
	"log"
	"time"
)

//todo 數據庫配置信息
const (
	DB_HOST    = "127.0.0.1"
	DB_PORT    = "3306"
	DB_USER    = "root"
	DB_PASS    = "recardo0126"
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
		log.Printf("[ERROR] 创建数据库连接失败: %s", err.Error())
		panic(err.Error())
	}
	db.LogMode(DB_DEBUG)

	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Users{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Chats{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Emails{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.IpUsers{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.Apps{})

	logs.Informational("连接数据库成功")

	//判断是否存在数据库记录
	user, count := make([]models.Users, 0), 0
	if err := db.Find(&user).Count(&count).Error; err != nil {
		panic("查询数据库失败！")
	} else if count == 0 {
		admin := new(models.Users)
		admin.User = "admin"
		admin.Pwd = utils.Sha1Pwd("123456")
		admin.Uuid = utils.NewUuid()
		admin.CreateAt = time.Now().Unix()
		if err := db.Create(admin).Error; err != nil {
			panic("创建初始用户失败")
		}
	}
}
