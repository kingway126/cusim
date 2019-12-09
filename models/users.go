package models

type Users struct {
	Id       int    `json:"id" gorm:" PRIMARY_KEY; AUTO_INCREMENT"`
	User     string `json:"user" gorm:"type:varchar(64); UNIQUE; NOT NULL"`
	Pwd      string `json:"pwd" gorm:"type:varchar(64); NOT NULL"`
	Hash     string `json:"hash" gorm:"type:varchar(64)"`
	Uuid     string `json:"uuid" gorm:"type:varchar(64); UNIQUE"`
	Email 	 string `json:"email" gorm:"type:varchar(124)"`
	CreateAt int64  `json:"create_at" gorm:"type:int(64)"`
}
