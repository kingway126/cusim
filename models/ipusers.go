package models

type IpUsers struct {
	Id       int    `json:"id" gorm:" PRIMARY_KEY; AUTO_INCREMENT"`
	Aid      int    `json:"aid" gorm:"type:int(10); NOT NULL"`
	Uid 	 int 	`json:"uid" gorm:"type:int(10); NOT NULL"`
	Ip       string `json:"ip" gorm:"type:varchar(64); NOT NULL"`
	Email    string  `json:"email, omitempty" gorm:"type:varchar(64)"`
	Name     string `json:"name, omitempty" gorm:"type:varchar(64)"`
	CreateAt int64 `json:"create_at, omitempty" gorm:"type:int(64)"`
	ConnectAt int64	`json:"connect_at, omitempty" gorm:"type:int(64)"`
}
