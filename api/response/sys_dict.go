package response

// 字典树信息响应,

type DictTreeResp struct {
	Id       uint           `json:"id"`
	ParentId uint           `json:"parent_id"`
	DictKey      string         `json:"dict_key"`
	DictValue    string         `json:"dict_value"`
	Desc     string         `json:"desc"`
	Sort     int            `json:"sort"`
	Creator  string         `json:"creator"`
	Children []DictTreeResp `json:"children,omitempty"` //tag:omitempty 为空的值不显示
}

type DictTreeRespList []DictTreeResp

func (hs DictTreeRespList) Len() int {
	return len(hs)
}
func (hs DictTreeRespList) Less(i, j int) bool {
	return hs[i].Sort < hs[j].Sort // 按Sort从小到大排序
}

func (hs DictTreeRespList) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}