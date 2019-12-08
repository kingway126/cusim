package utils
//登陆请求参数
type LoginArgs struct {
	User string `json:"user" form:"user"`
	Pass string `json:"pass" form:"pass"`
}
//
type TokenArgs struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
//resp
type LoginRespArgs struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
	Uuid  string 	`json:"uuid"`
}

//站点清单记录  请求参数格式
type AppArgs struct {
	TokenArgs
	PageSize  int    `json:"page_size"`
	PageIndex int    `json:"page_index"`
	Search    string `json:"search, omitempty"`
}
//删除站点
type DeleteSiteArgs struct {
	TokenArgs
	Aid 	int		`json:"aid"`
}
//更新站点信息
type SiteArgs struct {
	TokenArgs
	Aid 	int 	`json:"aid, omitempty"`
	Name 	string 	`json:"name"`
	Url 	string 	`json:"url"`
}
//定义chat的参数
type ChatArgs struct {
	Id 		int 	`json:"id"`
	Token 	string 	`json:"token, omitempty"`
	Uuid 	string 	`json:"uuid"`
	Role 	string 	`json："role, omitempty"`
}
