package controllers

type Msg struct {
	Srckey  string `json:"src_key"`
	DstKey  string `json:"dst_key"`
	SrcId   int    `json:"src_id"`
	DstId   int    `json:"dst_id"`
	SrcType int    `json:"src_type"`
	Type    int    `json:"type"`
	Data    string `json:"data"`
}

const (
	SRCTYPE_IP      = 0
	SRCTYPE_USER    = 1
	SRCTYPE_COMMAND = 2

	TYPE_MSG    = 0
	TYPE_FORM   = 1
	TYPE_NOTICE = 2
)
