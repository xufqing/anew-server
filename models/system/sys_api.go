package system

import "anew-server/models"

// 系统接口表
type SysApi struct {
	models.Model
	Name     string `gorm:"comment:'接口名称';size:64" json:"name"`
	Method   string `gorm:"comment:'请求方式';size:64" json:"method"`
	Path     string `gorm:"comment:'访问路径';size:128" json:"path"`
	PermsTag  string `gorm:"comment:'权限标识';size:128" json:"perms_tag"`
	Desc     string `gorm:"comment:'说明';size:128" json:"desc"`
	Creator  string `gorm:"comment:'创建人';size:64" json:"creator"`
	ParentId uint   `gorm:"default:0;comment:'父级接口(编号为0时表示根)'" json:"parent_id"`
}

func (m SysApi) TableName() string {
	return m.Model.TableName("sys_api")
}
