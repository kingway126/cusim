package utils

import "github.com/satori/go.uuid"

//todo 生成UUID
func NewUuid() string {
	u1 := uuid.Must(uuid.NewV4(), nil).String()
	return u1
}
