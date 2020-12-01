package request

import "anew-server/dto/response"

// 获取接口列表结构体
type HostListReq struct {
	HostName          string `json:"host_name" form:"host_name"`
	IpAddress         string `json:"ip_address" form:"ip_address"`
	OSVersion         string `json:"os_version" form:"os_version"`
	HostType          string `json:"host_type" form:"host_type"`
	AuthType          string `json:"auth_type" form:"auth_type"`
	Creator           string `json:"creator" form:"creator"`
	response.PageInfo        // 分页参数
}

// 创建接口结构体
type CreateHostReq struct {
	HostName  string `json:"host_name" form:"host_name"`
	IpAddress string `json:"ip_address" form:"ip_address" validate:"required"`
	HostType  string `json:"host_type" form:"host_type"`
	Prot      string `json:"prot" form:"prot"`
	AuthType  string `json:"auth_type" form:"auth_type" validate:"required"`
	User      string `json:"user" form:"user"`
	Password  string `json:"password" form:"password"`
	Creator   string `json:"creator" form:"creator"`
}

// 翻译需要校验的字段名称
func (s CreateHostReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["IpAddress"] = "主机地址"
	m["AuthType"] = "认证类型"
	return m
}
