package models

type Chats struct {
	Id       int    		`json:"id" gorm:" PRIMARY_KEY; AUTO_INCREMENT"`
	Iid      int    		`json:"iid" gorm:"type:int(10); NOT NULL"`
	Uid      int    		`json:"uid" gorm:"type:int(10); NOT NULL"`
	SrcType  string 		`json:"src_type" gorm:"type:varchar(4); NOT NULL"`
	Content  string 		`json:"content" gorm:"type:text; NOT NULL"`
	Read     string 		`json:"read" gorm:"type:varchar(4);"`
	CreateAt int64  		`json:"create_at" gorm:"type:int(64)"`
}

const (
	SRCTYPE_IP   = "ip"
	SRCTYPE_USER = "user"

	READ_YES = "yes"
	READ_NO  = "no"
)
