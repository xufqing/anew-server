package response

// 分组列表信息
type AssetGroupListResp struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
	Desc    string `json:"desc"`
	HostsId []uint `json:"hosts_id"`
}
