package request

import "anew-server/utils"

// 适用于大多数场景的请求参数绑定
type Req struct {
	Ids string `json:"ids" form:"ids"` // 传多个id
}

// 获取
func (s *Req) GetUintIds() []uint {
	return utils.Str2UintArr(s.Ids)
}


// 分页封装
type PageInfo struct {
	PageNum      uint `json:"pageNum" form:"pageNum"`           // 当前页码
	PageSize     uint `json:"pageSize" form:"pageSize"`         // 每页显示条数
	All bool `json:"all" form:"all"` // 不使用分页
}


// 计算limit/offset, 如果需要用到返回的PageSize, PageNum, 务必保证Total值有效
func (s *PageInfo) GetLimit() (limit uint, offset uint) {
	// 传入参数可能不合法, 设置默认值
	// 每页显示条数不能小于1
	if s.PageSize < 1 {
		s.PageSize = 10
	}
	// 页码不能小于1
	if s.PageNum < 1 {
		s.PageNum = 1
	}

	limit = s.PageSize
	offset = limit * (s.PageNum - 1)
	return
}