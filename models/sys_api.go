package models

// 系统接口表
type SysApi struct {
	Model
	Method   string `gorm:"comment:'请求方式';size:128" json:"method"`
	Path     string `gorm:"comment:'访问路径';size:255" json:"path"`
	Category string `gorm:"comment:'所属类别';size:128" json:"category"`
	Desc     string `gorm:"comment:'说明';size:255" json:"desc"`
	Creator  string `gorm:"comment:'创建人';size:128" json:"creator"`
	//Roles      []SysRole `gorm:"many2many:relation_role_api;" json:"roles"` // 角色接口多对多关系
}

func (m SysApi) TableName() string {
	return m.Model.TableName("sys_api")
}
