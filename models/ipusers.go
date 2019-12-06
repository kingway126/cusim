package models

type IpUsers struct {
	Id       int    `json:"id" gorm:" PRIMARY_KEY; AUTO_INCREMENT"`
	Aid      int    `json:"aid" gorm:"type:int(10)"`
	Ip       string `json:"ip" gorm:"type:varchar(64); UNIQUE"`
	Email    string `json:"email" gorm:"type:varchar(64)"`
	Name     string `json:"name" gorm:"type:varchar(64)"`
	CreateAt int64  `json:"create_at" gorm:"type:int(64)"`
}
