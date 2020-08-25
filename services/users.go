package services

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/recardoz/cusim/models"
	"github.com/recardoz/cusim/utils"
	"time"
)

//todo 根据用户名获取指定用户的信息
func GetUserInfo(user string) (*models.Users, error) {
	var users models.Users
	if err := db.Where("user = ?", user).First(&users).Error; err != nil {
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
func CheckUUID(user string) (string, error) {
	//查詢是否存在該用戶
	var users models.Users
	if err := db.Where("user = ?", user).First(&users).Error; users.Id == 0 {
		return "", errors.New("不存在該用戶")
	} else if err != nil {
		return "", err
	}
	if len(users.Uuid) == 0 {
		uuid := utils.NewUuid()
		createat := time.Now().Unix()
		if err := db.Model(&users).Updates(models.Users{Uuid: uuid, CreateAt: createat}).Error; err != nil {
			return "", err
		}
	}

	return users.Uuid, nil
}

//todo 通过uuid获取id
func CheckChat(id int, token, uuid string) (int, error) {
	user := new(models.Users)
	if err := db.Where("id = ? AND hash = ? AND uuid = ?", id, token, uuid).First(user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

//todo 更新用户的邮箱
func UpdateEmail(id int, email string) error {
	if err := db.Model(new(models.Users)).Where("id = ?", id).Update("email", email).Error; err != nil {
		return err
	}

	return nil
}

//todo 更新用户的密码
func UpdatePwd(id int, pwd string) error {
	if err := db.Model(new(models.Users)).Where("id = ?", id).Update("pwd", pwd).Error; err != nil {
		return err
	}

	return nil
}
