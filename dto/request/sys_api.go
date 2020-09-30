package request

import "anew-server/dto/response"

// 获取接口列表结构体
type ApiListReq struct {
	Method            string `json:"method" form:"method"`
	Path              string `json:"path" form:"path"`
	Category          string `json:"category" form:"category"`
	Creator           string `json:"creator" form:"creator"`
	Permission        string `json:"permission" form:"permission"`
	Tree              bool   `json:"tree" form:"tree"`
	response.PageInfo        // 分页参数
}

// 创建接口结构体
type CreateApiReq struct {
	Method   string `json:"method" validate:"required"`
	Path     string `json:"path" validate:"required"`
	Category string `json:"category" validate:"required"`
	Permission string `json:"permission" validate:"required"`
	Creator    string `json:"creator" form:"creator"`
	Desc     string `json:"desc"`
}

// 翻译需要校验的字段名称
func (s CreateApiReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Method"] = "请求方式"
	m["Path"] = "访问路径"
	m["Category"] = "所属类别"
	m["Permission"] = "权限标识"
	return m
}
