package response

import "anew-server/models"

// User login response structure
type LoginResp struct {
	Username  string `json:"username"`  // 登录用户名
	Token     string `json:"token"`     // jwt令牌
	ExpiresAt int64  `json:"expiresAt"` // 过期时间, 秒
}

// 用户信息响应
type UserInfoResp struct {
	Id          uint     `json:"id"`
	Username    string   `json:"username"`
	Mobile      string   `json:"mobile"`
	Avatar      string   `json:"avatar"`
	Name        string   `json:"name"`
	Email       string   `json:"mail"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

// 用户列表信息响应, 字段含义见models.SysUser
type UserListResp struct {
	Id        uint             `json:"id"`
	Username  string           `json:"username"`
	Mobile    string           `json:"mobile"`
	Avatar    string           `json:"avatar"`
	Name      string           `json:"name"`
	Email     string           `json:"mail"`
	Status    *bool            `json:"status"`
	RoleId    uint             `json:"roleId"`
	Creator   string           `json:"creator"`
	CreatedAt models.LocalTime `json:"createdAt"`
	UpdatedAt models.LocalTime `json:"updatedAt"`
}
