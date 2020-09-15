package models

// User
type SysUser struct {
	Model
	Username string  `gorm:"unique;comment:'用户名';size:128" json:"username" binding:"required"`
	Password string  `gorm:"comment:'密码';size:255" json:"password" binding:"required"`
	Mobile   string  `gorm:"comment:'手机';size:128" json:"mobile"`
	Avatar   string  `gorm:"comment:'头像';size:255" json:"avatar"`
	Name     string  `gorm:"comment:'姓名';size:128" json:"name"`
	Email    string  `gorm:"comment:'邮箱地址';size:128" json:"mail"`
	Status   *bool   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"` // 由于设置了默认值, 这里使用ptr, 可避免赋值失败
	Creator  string  `gorm:"comment:'创建人';size:128" json:"creator"`
	RoleId   uint    `gorm:"comment:'角色Id外键'" json:"roleId"`
	Role     SysRole `gorm:"foreignkey:RoleId" json:"role"` // 将SysUser.RoleId指定为外键
}

func (m SysUser) TableName() string {
	return m.Model.TableName("sys_user")
}
