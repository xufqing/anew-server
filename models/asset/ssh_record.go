package asset

import "anew-server/models"

// SSH审计表
type SSHRecord struct {
	models.Model
	ConnectID   string           `gorm:"comment:'连接标识';size:64" json:"connect_id"`
	UserName    string           `gorm:"comment:'系统用户名';size:128" json:"user_name"`
	HostName    string           `gorm:"comment:'主机名';size:128" json:"host_name"`
	IpAddress   string           `gorm:"comment:'主机地址';size:128" json:"ip_address"`
	ConnectTime models.LocalTime `gorm:"comment:'接入时间'" json:"connect_time"`
	LogoutTime  models.LocalTime `gorm:"comment:'注销时间'" json:"logout_time"`
	Records     string           `json:"records" gorm:"type:longblob;comment:'操作记录(二进制存储)';size:128"`
}

func (m SSHRecord) TableName() string {
	return m.Model.TableName("SSH_record")
}
