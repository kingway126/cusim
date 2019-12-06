package services

import (
	"testing"
	"github.com/astaxie/beego/logs"
	"strconv"
	"encoding/json"
	"CustomIM/models"
)

func clearAppsTable() {
	db.Exec("DELETE FROM apps")
	db.Exec("ALTER TABLE apps AUTO_INCREMENT = 1")
}
//todo 测试app相关操作
func TestAppsWorkFlow(t *testing.T) {
	clearAppsTable()
	t.Run("AddApp", testAddApp)
	t.Run("GetApp", testGetApp)
	t.Run("ReSetUuid", testResetUuid)
	t.Run("ReGetApp", testReGetApp)
	t.Run("DeleteApp", testDeleteApp)
	t.Run("ReGetDeleteApp", testReGetDeleteApp)
	t.Run("ListApp", testListApp)
	//clearAppsTable()
}
//增加app
func testAddApp(t *testing.T) {
	for i := 1; i <= 100; i ++ {
		err := NewApp(1,"大中华" + strconv.Itoa(i),"www.tc-docker.com","")
		if err != nil {
			t.Errorf("Error of AddApp: %v", err.Error())
		}
	}


}
//获取单个的记录
var appFirst *models.Apps
func testGetApp(t *testing.T) {
	var err error
	appFirst, err = GetAppInfo(1)
	if err != nil {
		t.Errorf("Error of GetApp: %v", err.Error())
	} else if appFirst.Name != "大中华1" || appFirst.Url != "www.tc-docker.com" || len(appFirst.Uuid) == 0 {
		t.Errorf("Error of GetApp: 获取记录失败")
	}
}
//重置uuid
func testResetUuid(t *testing.T) {
	err := ResetAppUuid(1)
	if err != nil {
		t.Errorf("Error of ResetAppUuid: %v", err.Error())
	}
}
//重新获取app信息，并比较Uuid
func testReGetApp(t *testing.T) {
	appSecond, err := GetAppInfo(1)
	if err != nil {
		t.Errorf("Error of ReGetApp: %v", err.Error())
	} else if appSecond.Uuid == appFirst.Uuid {
		t.Errorf("Error of ReGetApp: 重置uuid失败")
	}

}
//删除app记录
func testDeleteApp(t *testing.T) {
	err := DeleteApp(1)
	if err != nil {
		t.Errorf("Error of DeleteApp: %v", err.Error())
	}
}
//获取删除掉的app记录
func testReGetDeleteApp(t *testing.T) {
	appThird, err := GetAppInfo(1)
	if err != nil && err.Error() != "record not found" {
		t.Errorf("Error of ReGetDeleteApp: %v", err)
	} else if appThird != nil {
		t.Errorf("Error of ReGetDeleteApp: 删除记录失败")
	}
}
//获取多条app记录
func testListApp(t *testing.T) {
	all, apps, err := ListAppInfo(1,0,10,"")
	if	err != nil {
		t.Errorf("Error of ListApp: %s", err.Error())
	} else if all == 0 {
		t.Errorf("获取到0条记录")
	}
	logs.Informational("获取到", all, "条记录")
	for _, v := range apps {
		info , _ := json.Marshal(v)
		logs.Informational(string(info))
	}
}
