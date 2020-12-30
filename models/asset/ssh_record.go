package asset

import "anew-server/models"

// ssh审计表
type SshRecord struct {
	models.Model
	Key         string           `gorm:"comment:'标识';size:64" json:"key"`
	UserName    string           `gorm:"comment:'系统用户名';size:128" json:"user_name"`
	HostName    string           `gorm:"comment:'主机名';size:128" json:"host_name"`
	IpAddress   string           `gorm:"comment:'主机地址';size:128" json:"ip_address"`
	Port        string           `gorm:"comment:'Ssh端口';size:64" json:"port"`
	User        string           `gorm:"comment:'认证用户';size:64" json:"user"`
	ConnectTime models.LocalTime `gorm:"comment:'接入时间'" json:"connect_time"`
	LogoutTime  models.LocalTime `gorm:"comment:'注销时间'" json:"logout_time"`
	//InputData   string           `json:"input_data" gorm:"type:blob;comment:'输入命令(二进制存储)';size:128"`
	CastFileName    string           `json:"cast_file_name" gorm:"comment:'录像文件名';size:128"`
}

func (m SshRecord) TableName() string {
	return m.Model.TableName("ssh_record")
}
