package request

import (
	response2 "anew-server/api/response"
)

// 获取接口列表结构体
type HostReq struct {
	HostName           string `json:"host_name" form:"host_name"`
	IpAddress          string `json:"ip_address" form:"ip_address"`
	OSVersion          string `json:"os_version" form:"os_version"`
	HostType           string `json:"host_type" form:"host_type"`
	AuthType           string `json:"auth_type" form:"auth_type"`
	Creator            string `json:"creator" form:"creator"`
	GroupID            string `json:"group_id" form:"group_id"`
	response2.PageInfo        // 分页参数
}

// 创建接口结构体
type CreateHostReq struct {
	HostName      string `json:"host_name" form:"host_name"`
	IpAddress     string `json:"ip_address" form:"ip_address" validate:"required"`
	HostType      string `json:"host_type" form:"host_type"`
	Port          string `json:"port" form:"port"`
	AuthType      string `json:"auth_type" form:"auth_type" validate:"required"`
	User          string `json:"user" form:"user"`
	Password      string `json:"password" form:"password"`
	PrivateKey    string `json:"privatekey" form:"privatekey"`
	KeyPassphrase string `json:"key_passphrase"`
	Creator       string `json:"creator" form:"creator"`
}

// SSh结构体
type SShTunnelReq struct {
	HostId uint   `json:"hostId" form:"host_id"`
	Width  int    `json:"width" form:"width"`
	Hight  int    `json:"hight" form:"hight"`
	Token  string `json:"token" form:"token"`
}

// 文件管理req
type FileReq struct {
	HostId uint   `json:"host_id" form:"host_id"` // hostId
	Path   string `json:"path" form:"path"`
	Key    string `json:"key" form:"key"`
}

// 翻译需要校验的字段名称
func (s CreateHostReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["IpAddress"] = "主机地址"
	m["AuthType"] = "认证类型"
	return m
}
