package models

type Emails struct {
	Id       int    		`json:"id" gorm:" PRIMARY_KEY; AUTO_INCREMEN"`
	Host     string 		`json:"host" gorm:"type:varchar(64); NOT NULL"`
	Port     int    		`json:"port" gorm:"type:int(10); NOT NULL"`
	Secret   string 		`json:"secret" gorm:"type:varchar(64); NOT NULL"`
	Email    string 		`json:"email" gorm:type:"varchar(124); NOT NULL"`
	Name     string 		`json:"name" gorm:"type:varchar(64)"`
	CreateAt int64 	 	 	`json"create_at" gorm:"type:int(64)"`
}
