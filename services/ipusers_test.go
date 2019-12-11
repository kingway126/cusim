package services

import (
	"testing"
)

func clearTableIpUsers() {
	db.Exec("DELETE FROM ipusers")
	db.Exec("ALTER TABLE ipusers AUTO_INCREMENT=1")
}
//todo 测试Ipusers相关的操作
func TestIpUserWorkFlow(t *testing.T) {
	clearTableIpUsers()
	t.Run("FindOrCreate", testFindOrCreateIpUser)

}
//测试查询和创建ipuser的功能
func testFindOrCreateIpUser(t *testing.T) {
	for i := 1; i <= 3; i ++ {
		ipuser, err := FindOrCreateIpUser(3,1,"127.0.0.1")
		if err != nil {
			t.Errorf("Error of FindOrCreate: %v", err.Error())
		} else if ipuser.Id != 1 {
			t.Errorf("Error of FindOrCreate: 获取结果错误")
		}
	}

}
