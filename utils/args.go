package utils

type LoginArgs struct {
	User string `json:"user" form:"user"`
	Pass string `json:"pass" form:"pass"`
}

type TokenArgs struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

//站点清单记录  请求参数格式
type AppArgs struct {
	TokenArgs
	PageSize  int    `json:"page_size"`
	PageIndex int    `json:"page_index"`
	Search    string `json:"search, omitempty"`
}
