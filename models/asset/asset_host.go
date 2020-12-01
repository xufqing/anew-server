package asset

import "anew-server/models"

// 主机表
type AssetHost struct {
	models.Model
	HostName  string `gorm:"comment:'主机名';size:128" json:"host_name"`
	IpAddress string `gorm:"comment:'主机地址';size:128" json:"ip_address"`
	Prot      string    `gorm:"comment:'SSH端口';size:64" json:"prot"`
	OsVersion string `gorm:"comment:'系统版本';size:128" json:"os_version"`
	HostType  string `gorm:"comment:'主机类型';size:64" json:"host_type"`
	AuthType  string   `gorm:"comment:'认证类型'" json:"auth_type"`
	User  string   `gorm:"comment:'认证用户'" json:"user"`
	Password string `gorm:"comment:'认证密码';size:128" json:"password"`
	Creator   string `gorm:"comment:'创建人';size:64" json:"creator"`
}

func (m AssetHost) TableName() string {
	return m.Model.TableName("asset_host")
}
