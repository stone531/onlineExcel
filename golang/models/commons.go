package models

type TeamRole struct {
	TeamManagers
	Roles []Role `json:"roles"`
}

type Role struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	TeamId   int64  `json:"team_id"`
	RoleType string `json:"role_type"` // 角色类型：系统自动创建不允许编辑和删除 system,user
}

type TeamManagers struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	TeamType string `json:"team_type"` // 团队类型，默认为private,有一个公共团队为public
}

type LoginUser struct {
	UserName   string  // 中文名
	EmployeeId string  // 工号
	Account    string  // 英文名
	Email      string  // 邮箱
	TeamIds    []int64 // 所属团队
	TokenId    string
	IsAdmin    bool
	Teams      []TeamRole
	AuthType   string
}

//解析用户信息
func getLoginUser(userStr string) *LoginUser {
	loginUser := new(LoginUser)

	return loginUser
}

// 用户的 信息
type UserInfo struct {
	Username   string `json:"username"`   // 中文名
	Userstatus int64  `json:"userstatus"` //用户状态
	Mobile     string `json:"mobile"`
	//Employeeid string // 英文名
	Usercode string `json:"usercode"` // 工号
	Email    string `json:"email"`    // 邮箱
	Userid   string `json:"userid"`   //用户id
}

//请求参数
type UserInfoParam struct {
	Usercode string `json:"usercode"` // 工号
}

type JsonRet struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Err    string      `json:"error"`
	Data   interface{} `json:"data"`
}

type PermissionResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type AccountElement struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`
	Username   string `json:"username"`
	EmployeeId string `json:"employee_id"`
}
