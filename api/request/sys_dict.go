package request

// 创建字典结构体
type CreateDictReq struct {
	DictKey      string `json:"dict_key" validate:"required"`
	DictValue    string `json:"dict_value" validate:"required"`
	Sort     int    `json:"sort"`
	Desc     string `json:"desc"`
	ParentId uint   `json:"parent_id"`
	Creator  string `json:"creator"`
}

// 修改字典
type UpdateDictReq struct {
	DictKey      string `json:"dict_key" validate:"required"`
	DictValue    string `json:"dict_value" validate:"required"`
	Sort     int    `json:"sort"`
	Desc     string `json:"desc"`
	ParentId uint   `json:"parent_id"`
}

type DictListReq struct {
	DictKey     string `json:"dict_key" form:"key"`
	DictValue   string `json:"dict_value" form:"value"`
	Creator string `json:"creator" form:"creator"`
	TypeKey string `json:"type_key" form:"type_key"`
}

// 翻译需要校验的字段名称
func (s CreateDictReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["DictKey"] = "字典Key"
	m["DictValue"] = "字典Value"
	return m
}

// 翻译需要校验的字段名称
func (s UpdateDictReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["DictKey"] = "字典Key"
	m["DictValue"] = "字典Value"
	return m
}
