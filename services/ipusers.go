package services

import (
	"database/sql"
	"github.com/recardoz/cusim/models"
	"time"
)

type IpUserAndApp struct {
	models.IpUsers
	Email   sql.NullString `json:"email"`
	Name    sql.NullString `json:"name"`
	AppName string         `json:"app_name"`
	NoRead  sql.NullInt32  `json:"no_read"`
}

//todo 查询ip用户，如果找不到则创建该用户
func FindOrCreateIpUser(aid, uid int, ip string) (*models.IpUsers, error) {
	ipuser := models.IpUsers{
		Aid:       aid,
		Uid:       uid,
		Ip:        ip,
		CreateAt:  time.Now().Unix(),
		ConnectAt: time.Now().Unix(),
	}
	if err := db.Where("aid = ? AND uid = ? AND ip = ?", aid, uid, ip).FirstOrCreate(&ipuser).Error; err != nil {
		return nil, err
	}
	return &ipuser, nil
}

//todo 更新ip用户的链接时间
func UpdateConnect(iid int) error {
	ipuser := new(models.IpUsers)
	if err := db.Model(ipuser).Where("id = ?", iid).Update("connect_at", time.Now().Unix()).Error; err != nil {
		return err
	}
	return nil
}

//todo 通过ipuser的id，获取记录
func GetIpUserById(id int) (*models.IpUsers, error) {
	ipuser := new(models.IpUsers)
	if err := db.Where("id = ?", id).First(ipuser).Error; err != nil {
		return nil, err
	}
	return ipuser, nil
}

//todo 通过uid来查询指定connet_at时间内的用户
func ListIpUserForTime(uid int, begin, end int64) ([]*IpUserAndApp, error) {
	ipusers := make([]*IpUserAndApp, 0)
	rows, err := db.Raw(`SELECT a.*,c.name AS appname, b.noread FROM ip_users AS a LEFT JOIN apps AS c ON a.aid = c.id LEFT JOIN (SELECT iid, count(id) as noread FROM chats  WHERE uid = ? AND chats.isread = 'no' AND src_type = 'ip' GROUP BY iid)  AS b ON a.id = b.iid
 WHERE a.uid = ? AND a.connect_at > ? AND a.connect_at <= ? ORDER BY a.connect_at DESC`, uid, uid, begin, end).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := IpUserAndApp{}
		err := rows.Scan(&tmp.Id, &tmp.Aid, &tmp.Uid, &tmp.Ip, &tmp.Email, &tmp.Name, &tmp.CreateAt, &tmp.ConnectAt, &tmp.AppName, &tmp.NoRead)
		if err != nil {
			return ipusers, err
		}
		ipusers = append(ipusers, &tmp)
	}

	return ipusers, nil
}

//todo 查询ipuser的信息
func GetIpUser(iid, uid int) (*IpUserAndApp, error) {
	ipuser := new(IpUserAndApp)

	if err := db.Raw(`SELECT a.*, c.name as appname, b.noread FROM ip_users as a LEFT JOIN apps AS c ON a.aid = c.id LEFT JOIN (SELECT iid, COUNT(id) AS noread FROM chats WHERE iid = ? AND uid = ? GROUP BY iid) AS b ON a.id = b.iid WHERE a.id = ?`, iid, uid, iid).Scan(ipuser).Error; err != nil {
		return nil, err
	}
	return ipuser, nil
}

//todo 拆线呢ipuser的数量
func GetIpUserNum(uid int) (int, error) {
	var count int
	if err := db.Model(new(models.IpUsers)).Where("uid = ?", uid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
