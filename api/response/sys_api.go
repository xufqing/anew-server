package response

// 接口信息响应, 字段含义见models
type ApiListResp struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Creator  string `json:"creator"`
	Desc     string `json:"desc"`
	PermsTag string `json:"perms_tag"`
	//CreatedAt models.LocalTime `json:"created_at"`
	ParentId uint          `json:"parent_id"`
	Children []ApiListResp `json:"children,omitempty"`
}
