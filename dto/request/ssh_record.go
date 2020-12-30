package request

import (
	"anew-server/dto/response"
)

// 获取操作日志列表结构体
type SshRecordListReq struct {
	Key               string           `json:"key" form:"key"`
	UserName          string           `json:"user_name" form:"user_name"`
	HostName          string           `json:"host_name" form:"host_name"`
	IpAddress         string           `json:"ip_address" form:"ip_address"`
	response.PageInfo                  // 分页参数
}
