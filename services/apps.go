package services

import (
	"CustomIM/models"
	"CustomIM/utils"
	"github.com/astaxie/beego/logs"
	"time"
	"errors"
)

//todo 创建app
func NewApp(uid int, name, url, icon string) error {
	uuid := utils.NewUuid()
	app := models.Apps{
		Uid: 	uid,
		Name: 	name,
		Url: 	url,
		Icon:	icon,
		Uuid: 	uuid,
		CreateAt: 	time.Now().Unix(),
	}

	//新增记录
	if err := db.Create(&app).Error; err != nil {
		logs.Informational(err.Error())
		return err
	} else if app.Id == 0 {
		logs.Informational("添加站点失败")
		return errors.New("添加站点失败")
	}

	return nil
}
//todo 获取单个的app信息
func GetAppInfo(id int) (*models.Apps, error) {
	app := new(models.Apps)
	if err := db.Where("id = ?", id).First(&app).Error; err != nil {
		logs.Informational(err.Error())
		return nil, err
	}
	return app, nil
}
//todo 获取多个的app信息
func ListAppInfo(id, pageindex, pagesize int, search string) (int, []*models.Apps, error) {
	app := new(models.Apps)
	apps := make([]*models.Apps, 0)
	var all int

	//获取多条信息
	rows, err := db.Model(app).Where("uid = ? AND ( name LIKE ? OR  url LIKE ?)", id, "%" + search + "%", "%" + search + "%").Count(&all).Offset(pageindex).Limit(pagesize).Rows()
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := new(models.Apps)
		err := rows.Scan(&tmp.Id,&tmp.Uid, &tmp.Name, &tmp.Url, &tmp.Icon, &tmp.Uuid, &tmp.CreateAt)
		if err != nil {
			return all, apps, err
		}
		apps = append(apps, tmp)
	}

	return all, apps, nil
}
//todo 删除某条记录
func DeleteApp(id int) error {
	app := models.Apps{Id: id}
	if err := db.Delete(&app).Error; err != nil {
		logs.Informational(err.Error())
		return err
	}
	return nil
}
//todo 重置某个app的uuid
func ResetAppUuid(id int) error {
	uuid := utils.NewUuid()
	if err := db.Model(new(models.Apps)).Where("id = ?", id).Update("uuid", uuid).Error; err != nil {
		logs.Informational(err.Error())
		return err
	}
	return nil
}
//todo 修改App的信息
func ChangeApp(id int, name, url, icon string) error {
	if err := db.Model(new(models.Apps)).Where("id = ?", id).Updates(models.Apps{Name: name, Url: url, Icon: icon}).Error; err != nil {
		logs.Informational(err.Error())
		return err
	}
	return nil
}