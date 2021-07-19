package request

import (
	response "anew-server/api/response"
)

// 获取操作日志列表结构体
type SSHRecordReq struct {
	UserName          string `json:"user_name" form:"user_name"`
	response.PageInfo        // 分页参数
}
