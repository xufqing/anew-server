package request

import (
	"anew-server/models"
	"anew-server/utils"
)

// 适用于前端传过来的
type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"` // 传多个id
}

// 增量更新id集合结构体
type UpdateIncrementalIdsReq struct {
	Create []uint `json:"create"` // 需要新增的编号集合
	Delete []uint `json:"delete"` // 需要删除的编号集合
}

// 获取增量, 可直接更新的结果
func (s *UpdateIncrementalIdsReq) GetIncremental(oldMenuIds []uint, allMenu []models.SysMenu) []uint {
	// 保留选中流水线
	s.Create = models.GetCheckedMenuIds(s.Create, allMenu)
	s.Delete = models.GetCheckedMenuIds(s.Delete, allMenu)
	newList := make([]uint, 0)
	for _, oldItem := range oldMenuIds {
		// 已删除数据不加入列表
		if !utils.Contains(s.Delete, oldItem) {
			newList = append(newList, oldItem)
		}
	}
	// 将需要新增的数据合并
	return append(newList, s.Create...)
}
