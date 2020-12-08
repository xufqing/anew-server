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
	UserName       string           `json:"username"`
	Name           string           `json:"name"`
	HostName       string           `json:"host_name"`
	IpAddress      string           `json:"ip_address"`
	Port           string           `json:"port"`
	ConnectionTime models.LocalTime `json:"connection_time"`
	LastActiveTime models.LocalTime `json:"last_active_time"`
}