package system

import "anew-server/models"

// 系统部门表
type SysDept struct {
	models.Model
	Name     string    `gorm:"comment:'部门名称';size:64" json:"name"`
	Sort     int       `gorm:"type:int(3);comment:'排序'" json:"sort"`
	ParentId uint      `gorm:"default:0;comment:'父级部门(编号为0时表示根)'" json:"parent_id"`
	Creator  string    `gorm:"comment:'创建人';size:128" json:"creator"`
	Children []SysDept `gorm:"-" json:"children"` // 下属部门集合
	Users    []SysUser `gorm:"foreignkey:DeptId"` // 一个部门有多个user
}

func (m SysDept) TableName() string {
	return m.Model.TableName("sys_dept")
}
