package utils

const (
	//消息的来源
	SRCTYPE_IP      = "ip"
	SRCTYPE_USER    = "user"
	SRCTYPE_ORDER   = "order"
	//消息的类型
	CMD_MSG    = "msg"
	CMD_FORM   = "form"
	CMD_NOTICE = "notice"
	//节点对应的key的前缀
	PREFIX_USER = "admin"
	PREFIX_IP 	 = "ip"
)
//消息体结构
type Msg struct {
	DstKey  string 			`json:"dst_key"`
	SrcId   int    			`json:"src_id, omitempty"`
	DstId   int    			`json:"dst_id"`
	SrcType string 			`json:"src_type"`
	Cmd     string 			`json:"cmd"`
	Data    interface{}	`json:"data"`
}