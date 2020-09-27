package response

import (
	"anew-server/models"
)

// 角色返回菜单信息
type RoleMenusResp struct {
	Id       uint            `json:"id"`
	Name     string          `json:"name"`
	ParentId uint            `json:"parentId"`
	Children []RoleMenusResp `json:"children"`
	Status   *bool           `json:"status"`
}

// 角色信息响应, 字段含义见models
type RoleListResp struct {
	Id        uint             `json:"id"`
	Key       string           `json:"key"`
	Name      string           `json:"name"`
	Title     string           `json:"title"`
	Keyword   string           `json:"keyword"`
	Desc      string           `json:"desc"`
	Status    *bool            `json:"status"`
	Menus     []RoleMenusResp  `json:"menus"`
	Apis      []string         `json:"apis"`
	Creator   string           `json:"creator"`
	CreatedAt models.LocalTime `json:"createdAt"`
}
