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
	GroupID  	int  			`json:"group_id"`
	IpId   	int    				`json:"ip_id"`
	SrcType 	string 			`json:"src_type"`
	Cmd     	string 			`json:"cmd"`
	Data    	string 			`json:"content"`
	Date 		int64 			`json:"create_at"`
	Iid 		int 			`json:"iid, omitempty"`
}

