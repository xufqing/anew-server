package request

// 创建部门结构体
type CreateDeptReq struct {
	Name      string `json:"name" validate:"required"`
	Sort       int    `json:"sort"`
	Status     *bool  `json:"status"`
	ParentId   uint   `json:"parentId"`
	Creator    string `json:"creator"`
}

type DeptListReq struct {
	Name              string `json:"name" form:"name"`
	Status            *bool  `json:"status" form:"status"`
}

// 翻译需要校验的字段名称
func (s CreateDeptReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Name"] = "部门名称"
	return m
}
