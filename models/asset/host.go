package asset

import "anew-server/models"

// 主机表
type AssetHost struct {
	models.Model
	HostName     string `gorm:"comment:'主机名';size:128" json:"host_name"`
	IpAddress   string `gorm:"comment:'主机地址';size:128" json:"ip_address"`
	OsVersion     string `gorm:"comment:'系统版本';size:128" json:"os_version"`
	HostType     string `gorm:"comment:'主机类型';size:128" json:"host_type"`
	AuthType     uint `gorm:"comment:'认证类型'" json:"auth_type"`
	Creator  string `gorm:"comment:'创建人';size:128" json:"creator"`
}

func (m AssetHost) TableName() string {
	return m.Model.TableName("asset_host")
}
