package response

import (
	"anew-server/models"
)

type HostListResp struct {
	Id         uint   `json:"id"`
	HostName   string `json:"host_name"`
	IpAddress  string `json:"ip_address"`
	Port       string `json:"port"`
	OsVersion  string `json:"os_version"`
	HostType   string `json:"host_type"`
	AuthType   string `json:"auth_type"`
	User       string `json:"user"`
	PrivateKey string `json:"privatekey"`
	Creator    string `json:"creator"`
}

type SessionResp struct {
	ConnectID   string           `json:"connect_id"`
	UserName    string           `json:"user_name"`
	HostName    string           `json:"host_name"`
	IpAddress   string           `json:"ip_address"`
	ConnectTime models.LocalTime `json:"connect_time"`
}

type SessionRespList []SessionResp

func (hs SessionRespList) Len() int {
	return len(hs)
}
func (hs SessionRespList) Less(i, j int) bool {

	return hs[i].ConnectTime.Time.Before(hs[j].ConnectTime.Time) // 按Sort从小到大排序
}

func (hs SessionRespList) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

type SSHRecordListResp struct {
	ConnectID   string           `json:"connect_id"`
	UserName    string           `json:"user_name"`
	HostName    string           `json:"host_name"`
	IpAddress   string           `json:"ip_address"`
	ConnectTime models.LocalTime `json:"connect_time"`
	LogoutTime  models.LocalTime `json:"logout_time"`
}

type FileInfo struct {
	Name   string           `json:"name"`
	Path   string           `json:"path"`
	IsDir  bool             `json:"isDir"`
	Mode   string           `json:"mode"`
	Size   string           `json:"size"`
	Mtime  models.LocalTime `json:"mtime"` // 修改时间
	IsLink bool             `json:"isLink"`
}
