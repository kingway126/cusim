package services

import (
	"testing"
	"CustomIM/models"
)

func clearTableChats() {
	db.Exec("DELETE FROM chats")
	db.Exec("ALTER TABLE chats AUTO_INCREMENT=1")
}
//todo chats表的相关操作
func TestChatsWorkFlow(t *testing.T) {
	clearTableChats()
	t.Run("AddChatMsg", testAddChatMsg)
	t.Run("ListChatMsg", testListChatMsg)
}
// 添加聊天记录
func testAddChatMsg(t *testing.T) {
	err := AddChatMsg(1,1,1, "ip", "你好", models.READ_NO, 123234)
	if err != nil {
		t.Errorf("Error of AddChatMsg: %v", err.Error())
	}
}
// 获取所有的聊天记录
func testListChatMsg(t *testing.T) {
	chats, err := ListChatMsg(1, 1, 1)
	if err != nil {
		t.Errorf("Error of ListChatMsg: %v", err.Error())
	} else if len(chats) != 1 {
		t.Errorf("Error of ListChatMsg: Wrong length")
	}

}
