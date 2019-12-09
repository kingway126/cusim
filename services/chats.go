package services

import "CustomIM/models"
//todo 添加新的聊天记录
func AddChatMsg(aid, iid, uid int, srctype string, content string, read string, createat int64) error {
	chat := models.Chats{
		Aid: 		aid,
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
func ListChatMsg(aid, iid, uid int) ([]*models.Chats, error) {
	res := make([]*models.Chats, 0)
	rows, err := db.Model(new(models.Chats)).Where("aid = ? AND iid = ? AND uid = ?", aid, iid, uid).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := new(models.Chats)
		err := rows.Scan(&tmp.Id, &tmp.Aid, &tmp.Iid, &tmp.Uid, &tmp.SrcType, &tmp.Content, &tmp.Read, &tmp.CreateAt)
		if err != nil {
			return res, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
