package asset

import "anew-server/models"

// 分组
type AssetGroup struct {
	models.Model
	Name    string      `gorm:"comment:'分组名称';size:128" json:"name"`
	Creator string      `gorm:"comment:'创建人';size:64" json:"creator"`
	Desc    string    `gorm:"comment:'分组说明';size:255" json:"desc"`
	Hosts   []AssetHost `gorm:"many2many:relation_group_host;" json:"hosts"`
}

func (m AssetGroup) TableName() string {
	return m.Model.TableName("asset_group")
}
