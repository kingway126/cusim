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
//更新用户邮箱
type UserInfo struct {
	TokenArgs
	Email 	string	`json:"email"`
}
//更新用户密码
type UserPwd struct {
	TokenArgs
	Pwd 	string 	`json:"pwd"`
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
//加载聊天记录，Iid
type ChatListArgs struct {
	TokenArgs
	Iid 	int
}
//首页数据
type IndexNum struct {
	App 	int		`json:"app"`
	User 	int		`json:"user"`
	NoRead	int		`json:"noread"`
	Read 	int		`json:"read"`
}
