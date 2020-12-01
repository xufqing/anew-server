package response

type HostListResp struct {
	Id        uint   `json:"id"`
	HostName  string `json:"host_name"`
	IpAddress string `json:"ip_address"`
	Prot      string `json:"prot"`
	OsVersion string `json:"os_version"`
	HostType  string `json:"host_type"`
	AuthType  string `json:"auth_type"`
	User      string `json:"user"`
	Creator   string `json:"creator"`
}
