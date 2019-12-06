package services

import (
	"CustomIM/models"
	"CustomIM/utils"
	"database/sql"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"time"
)

//todo 根据用户名获取指定用户的信息
func GetUserInfo(user string) (*models.Users, error) {
	var users models.Users
	if err := db.Where("user = ?", user).First(&users).Error; err != nil {
		logs.Informational(err.Error())
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, errors.New("不存在該用戶")
	}
	return &users, nil
}

//todo 根据用户ID获取指定用户的信息
func GetUserInfoById(id int) (*models.Users, error) {
	var users models.Users
	if err := db.Where("id = ?", id).First(&users).Error; err != nil {
		logs.Informational(err.Error())
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, errors.New("不存在該用戶")
	}
	return &users, nil
}

//todo 更新用户token
func UpdateUserToken(user string) (string, error) {
	//查詢是否存在該用戶
	var users models.Users
	if err := db.Where("user = ?", user).First(&users).Error; users.Id == 0 {
		return "", errors.New("不存在該用戶")
	} else if err != nil {
		return "", err
	}

	//更新token和createat
	token := utils.Sha1Token(user)
	createat := time.Now().Unix()

	if err := db.Model(&users).Updates(models.Users{Hash: token, CreateAt: createat}).Error; err != nil {
		return "", err
	}

	return users.Hash, nil
}

//todo 检测用户uuid，如果没有，自动创建
func CheckUUID(user string) error {
	//查詢是否存在該用戶
	var users models.Users
	if err := db.Where("user = ?", user).First(&users).Error; users.Id == 0 {
		return errors.New("不存在該用戶")
	} else if err != nil {
		return err
	}
	uuid := utils.NewUuid()
	createat := time.Now().Unix()
	if err := db.Model(&users).Updates(models.Users{Uuid: uuid, CreateAt: createat}).Error; err != nil {
		return err
	}

	return nil
}
