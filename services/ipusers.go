package services

import (
	"CustomIM/models"
	"time"
)

//todo 查询ip用户，如果找不到则创建该用户
func FindOrCreateIpUser(aid, uid int, ip string ) (*models.IpUsers, error) {
	ipuser := models.IpUsers{
		Aid: 		aid,
		Uid: 		uid,
		Ip: 		ip,
		CreateAt: 	time.Now().Unix(),
		ConnectAt: 	time.Now().Unix(),
	}
	if err := db.Where("aid = ? AND uid = ? AND ip = ?", aid, uid, ip).FirstOrCreate(&ipuser).Error; err != nil {
		return nil, err
	}
	return &ipuser, nil
}
//todo 通过ipuser的id，获取记录
func GetIpUserById(id int) (*models.IpUsers, error) {
	ipuser := new(models.IpUsers)
	if err := db.Where("id = ?", id).First(ipuser).Error; err != nil {
		return nil, err
	}
	return ipuser, nil
}
