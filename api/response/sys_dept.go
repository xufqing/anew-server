package response

// 部门树信息响应,
type DeptTreeResp struct {
	Id       uint           `json:"id"`
	ParentId uint           `json:"parent_id"`
	Name     string         `json:"name"`
	Creator  string         `json:"creator"`
	Sort     int            `json:"sort"`
	Children []DeptTreeResp `json:"children,omitempty"` //tag:omitempty 为空的值不显示
}

type DeptTreeRespList []DeptTreeResp

func (hs DeptTreeRespList) Len() int {
	return len(hs)
}
func (hs DeptTreeRespList) Less(i, j int) bool {
	return hs[i].Sort < hs[j].Sort // 按Sort从小到大排序
}

func (hs DeptTreeRespList) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
