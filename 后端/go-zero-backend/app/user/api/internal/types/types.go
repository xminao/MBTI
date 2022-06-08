// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Username         string `json:"username"`           // 用户名
	Nickname         string `json:"nickname"`           // 昵称
	Password         string `json:"password"`           // 密码
	Gender           string `json:"gender"`             // 性别
	AuthGroup        string `json:"auth_group"`         // 认证权限
	BindingStudentId string `json:"binding_student_id"` // 绑定的学号
	CreatedAt        int64  `json:"created_at"`         // 创建时间
	UpdateAt         int64  `json:"update_at"`          // 更新时间
}

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type RegisterReq struct {
	Username         string `json:"username"`
	NickName         string `json:"nickname"`
	Password         string `json:"password"`
	Gender           string `json:"gender"`
	BindingStudentId string `json:"binding_student_id"`
}

type RegisterResp struct {
	Msg      string `json:"msg"`
	Nickname string `json:"nickname"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Msg      string   `json:"msg"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Jwt      JwtToken `json:"jwt_token"`
}

type GetUserInfoReq struct {
	Username string `form:"username"`
}

type GetUserInfoResp struct {
	UserInfoResp UserInfo `json:"userinfo"`
}

type GetUserListReq struct {
}

type GetUserListResp struct {
	UserList []UserInfo `json:"userlist"`
}

type VerifyTokenReq struct {
}

type VerifyTokenResp struct {
	Username string `json:"username"`
}
