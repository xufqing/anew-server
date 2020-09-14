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

// 获取用户列表结构体
type UserListReq struct {
	Id       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string `json:"avatar" form:"avatar"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"mail" form:"mail"`
	Status   *bool  `json:"status" form:"status"`
	RoleId   uint   `json:"roleId" form:"roleId"`
	Creator  string `json:"creator" form:"creator"`
	PageInfo        // 分页参数
}

// 创建用户结构体
type CreateUserReq struct {
	Username     string `json:"username" binding:"required"`
	InitPassword string `json:"initPassword" binding:"required"` // 不使用SysUser的Password字段, 避免请求劫持绕过系统校验
	Mobile       string `json:"mobile" binding:"len=11"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"mail"`
	Status       *bool  `json:"status"`
	RoleId       uint   `json:"roleId" binding:"required"`
	Creator      string `json:"creator"`
}
