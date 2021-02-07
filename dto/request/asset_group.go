package request

import (
	"anew-server/dto/response"
)

type AssetGroupReq struct {
	Name              string `json:"name" form:"name"`
	Creator           string `json:"creator" form:"creator"`
	response.PageInfo        // 分页参数
}

type CreateAssetGroupReq struct {
	Name    string `json:"name" validate:"required"`
	Creator string `json:"creator"`
	Desc      string           `json:"desc"`
	Hosts   []uint `json:"hosts"`
}

// 修改分组
type UpdateAssetGroupReq struct {
	Name   string `json:"name" validate:"required"`
	Desc      string           `json:"desc"`
	Hosts  []uint `json:"hosts"`
}

// 翻译需要校验的字段名称
func (s CreateAssetGroupReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Name"] = "分组名称"
	return m
}
