package system

import "anew-server/models"

// 数据字典
type SysDict struct {
	models.Model
	Key      string    `gorm:"uniqueIndex:uk_key;comment:'字典Key';size:64" json:"key"`
	Value    string    `gorm:"comment:'字典Value';size:64" json:"value"`
	Desc     string    `gorm:"comment:'说明';size:128" json:"desc"`
	Creator  string    `gorm:"comment:'创建人';size:64" json:"creator"`
	ParentId uint      `gorm:"default:0;comment:'父级字典(编号为0时表示根)'" json:"parent_id"`
	Dicts    []SysDict `gorm:"foreignkey:ParentId"`
}

func (m SysDict) TableName() string {
	return m.Model.TableName("sys_dict")
}
