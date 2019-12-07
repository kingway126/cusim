package controllers

type Msg struct {
	Srckey  string `json:"src_key"`
	DstKey  string `json:"dst_key"`
	SrcId   int    `json:"src_id"`
	DstId   int    `json:"dst_id"`
	SrcType string `json:"src_type"`
	Type    string `json:"type"`
	Data    string `json:"data"`
}

const (
	SRCTYPE_IP      = 0x00
	SRCTYPE_USER    = 0x01
	SRCTYPE_COMMAND = 0x02

	TYPE_MSG    = 0x00
	TYPE_FORM   = 0x01
	TYPE_NOTICE = 0x02
)
