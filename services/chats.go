package services

import "CustomIM/models"
//todo 添加新的聊天记录
func AddChatMsg(iid, uid int, srctype string, content string, read string, createat int64) error {
	chat := models.Chats{
		Iid: 		iid,
		Uid: 		uid,
		SrcType:	srctype,
		Content: 	content,
		Read: 		read,
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
	rows, err := db.Model(new(models.Chats)).Where("iid = ? AND uid = ?", iid, uid).Update("read", "yes").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		tmp := new(models.Chats)
		err := rows.Scan(&tmp.Id, &tmp.Iid, &tmp.Uid, &tmp.SrcType, &tmp.Content, &tmp.Read, &tmp.CreateAt)
		if err != nil {
			return res, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
//todo 将某个ipuser的所以聊天记录设置成已读
func ChatReadHad(iid int) error {
	if err := db.Model(new(models.Chats)).Where("id = ?", iid).Update("read", "yes").Error; err != nil {
		return err
	}
	return nil
}
