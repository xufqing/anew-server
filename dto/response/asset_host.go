package response

import "anew-server/models"

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

type ConnectionResp struct {
	Key            string           `json:"key"`
	UserName       string           `json:"user_name"`
	Name           string           `json:"name"`
	HostName       string           `json:"host_name"`
	IpAddress      string           `json:"ip_address"`
	Port           string           `json:"port"`
	ConnectTime models.LocalTime `json:"connect_time"`
}

type FileInfo struct {
	Name   string           `json:"name"`
	Path   string           `json:"path"`
	IsDir  bool             `json:"isDir"`
	Mode   string           `json:"mode"`
	Size   string            `json:"size"`
	Mtime  models.LocalTime `json:"mtime"` // 修改时间
	IsLink bool             `json:"isLink"`
}
