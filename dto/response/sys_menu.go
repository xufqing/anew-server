package response

// 菜单树信息响应, 字段含义见models.SysMenu
type MenuTreeResp struct {
	Id         uint           `json:"id"`
	ParentId   uint           `json:"parentId"`
	Name       string         `json:"name"`
	Title      string         `json:"title"`
	Icon       string         `json:"icon"`
	Path       string         `json:"path"`
	Creator    string         `json:"creator"`
	Sort       int            `json:"sort"`
	Status     bool           `json:"status"`
	Hidden    bool           `json:"hidden"`
	Children   []MenuTreeResp `json:"children"`
}

// 菜单树信息响应, 包含有权限访问的id列表
type MenuTreeWithAccessResp struct {
	List      []MenuTreeResp `json:"list"`
	AccessIds []uint         `json:"accessIds"`
}
