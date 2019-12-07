package models

type Apps struct {
	Id       int    `json:"id" gorm:" PRIMARY_KEY; AUTO_INCREMENT"`
	Uid      int    `json:"uid" gorm:"type:int(10); NOT NULL"`
	Name     string `json:"name" gorm:"type:varchar(64); NOT NULL UNIQUE"`
	Url      string `json:"url" gorm:"type:varchar(64)"`
	Icon     string `json:"icon" gorm:"type:varchar(255)"`
	Uuid     string `json:"uuid" gorm:"type:varchar(64); UNIQUE"`
	CreateAt int64  `json:"create_at" gorm:"type:int(64)"`
}
