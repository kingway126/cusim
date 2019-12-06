package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"time"
)

//todo 加密用户密码
func Sha1Pwd(pwd string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(pwd))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

//todo 计算用户token
func Sha1Token(user string) string {
	t := strconv.Itoa(int(time.Now().Unix()))
	sha1 := sha1.New()
	sha1.Write([]byte(user + t))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}
