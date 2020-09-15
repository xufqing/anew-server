package request

// User login structure
type RegisterAndLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 修改密码结构体
type ChangePwdReq struct {
	OldPassword string `json:"oldPassword" form:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" validate:"required"`
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
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"` // 不使用SysUser的Password字段, 避免请求劫持绕过系统校验
	Mobile   string `json:"mobile" validate:"len=11"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"mail"`
	Status   *bool  `json:"status"`
	RoleId   uint   `json:"roleId" validate:"required"`
	Creator  string `json:"creator"`
}

// 翻译需要校验的字段名称
func (s CreateUserReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["Password"] = "用户密码"
	m["Mobile"] = "手机号码"
	m["Name"] = "姓名"
	m["RoleId"] = "角色ID"
	return m
}
func (s ChangePwdReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["OldPassword"] = "旧密码"
	m["NewPassword"] = "新密码"
	return m
}
