package request

// User login structure
type RegisterAndLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 修改密码结构体
type ChangePwdReq struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

// 创建用户结构体
type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"` // 不使用SysUser的Password字段, 避免请求劫持绕过系统校验
	Mobile   string `json:"mobile" validate:"eq=11"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	Email    string `json:"mail"`
	Status   *bool  `json:"status"`
	RoleId   uint   `json:"roleId" validate:"required"`
	Creator  string `json:"creator"`
}
