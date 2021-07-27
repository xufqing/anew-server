package request

// 创建接口结构体
type CreateApiReq struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	PermsTag string `json:"perms_tag" validate:"required"`
	Creator  string `json:"creator" form:"creator"`
	ParentId uint   `json:"parent_id"`
	Desc     string `json:"desc"`
}

// 修改接口结构体
type UpdateApiReq struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	PermsTag string `json:"perms_tag" validate:"required"`
	ParentId uint   `json:"parent_id"`
	Desc     string `json:"desc"`
}

// 翻译需要校验的字段名称
func (s CreateApiReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Name"] = "接口名称"
	m["PermsTag"] = "权限标识"
	return m
}
