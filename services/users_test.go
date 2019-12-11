package services

import (
	"CustomIM/utils"
	"fmt"
	"testing"
)

var token string

//todo 测试用户相关操作
func TestUserFlow(t *testing.T) {
	t.Run("GetUser", testGetUser)
	t.Run("UpdateToken", testUpdateToken)
	t.Run("CheckToken", testCheckToken)
}

func testGetUser(t *testing.T) {
	_, err := GetUserInfo("admin")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err.Error())
	}
}

func testUpdateToken(t *testing.T) {
	var err error
	token, err = UpdateUserToken("admin")
	fmt.Println("token", utils.Sha1Token("admin"))
	if err != nil {
		t.Errorf("Error of UpdateToken: %v", err.Error())
	}
}

func testCheckToken(t *testing.T) {
	user, err := GetUserInfo("admin")
	if err != nil {
		t.Errorf("Error of CheckToken: %v", err.Error())
	} else if user.Hash != token {
		t.Errorf("Error of CheckToken: token錯誤")
	}
}
