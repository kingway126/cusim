package services

import "CustomIM/models"
//todo 添加新的聊天记录
func AddChatMsg(iid, uid int, srctype string, content string, read string, createat int64) error {
	chat := models.Chats{
		Iid: 		iid,
		Uid: 		uid,
		SrcType:	srctype,
		Content: 	content,
		Isread: 		read,
		CreateAt: 	createat,
	}
	if err := db.Create(&chat).Error; err != nil {
		return err
	}

	return nil
}
//todo 获取所有的记录
func ListChatMsg(iid, uid int) ([]*models.Chats, error) {
	res := make([]*models.Chats, 0)
	rows, err := db.Model(new(models.Chats)).Where("iid = ? AND uid = ?", iid, uid).Update("isread", "yes").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		tmp := new(models.Chats)
		err := rows.Scan(&tmp.Id, &tmp.Iid, &tmp.Uid, &tmp.SrcType, &tmp.Content, &tmp.Isread, &tmp.CreateAt)
		if err != nil {
			return res, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
//todo 将某个ipuser的所以聊天记录设置成已读
func ChatReadHad(iid int) error {
	if err := db.Model(new(models.Chats)).Where("id = ?", iid).Update("isread", "yes").Error; err != nil {
		return err
	}
	return nil
}
//todo 获取所有未读的数量
func GetNoReadNum(uid int, read string) (int, error) {
	var count int
	if err := db.Model(new(models.Chats)).Where("uid = ? AND isread = ?", uid, read).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
//todo 获取所有的数量
func GetAllNum(uid int) (int, error) {
	var count int
	if err := db.Model(new(models.Chats)).Where("uid = ?", uid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
//todo 将消息设置成已读
func SetReadAll(uid, iid int) error {
	if err := db.Model(new(models.Chats)).Where("uid = ? AND iid = ?", uid, iid).Update("isread", models.READ_YES).Error; err != nil {
		return err
	}
	return  nil
}
