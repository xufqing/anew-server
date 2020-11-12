package response

import (
	"anew-server/models"
)

// 接口信息响应, 字段含义见models
type ApiListResp struct {
	Id         uint             `json:"id"`
	Key        string           `json:"key"`
	Method     string           `json:"method"`
	Path       string           `json:"path"`
	Category   string           `json:"category"`
	Creator    string           `json:"creator"`
	Desc       string           `json:"desc"`
	Title      string           `json:"title"`
	CreatedAt  models.LocalTime `json:"created_at"`
}

type ApiTreeResp struct {
	Key      string        `json:"key"`
	Title    string        `json:"title"`    // 标题
	Category string        `json:"category"` // 分组名称
	Children []ApiListResp `json:"children"` // 前端以树形图结构展示, 这里用children表示
}
