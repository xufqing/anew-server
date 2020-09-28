package request

// 适用于前端传过来的
type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"` // 传多个id
}
