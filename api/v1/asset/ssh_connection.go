package asset

import (
	request2 "anew-server/api/request"
	response2 "anew-server/api/response"
	"github.com/gin-gonic/gin"
)

// 获取连接列表
func GetConnections(c *gin.Context) {
	var resp []response2.ConnectionResp
	for client, _ := range hub.Clients {
		var connStruct response2.ConnectionResp
		connStruct.Key = hub.Clients[client].Key
		connStruct.UserName = hub.Clients[client].UserName
		connStruct.Name = hub.Clients[client].Name
		connStruct.HostName = hub.Clients[client].HostName
		connStruct.IpAddress = hub.Clients[client].IpAddress
		connStruct.Port = hub.Clients[client].Port
		connStruct.ConnectTime = hub.Clients[client].ConnectTime
		resp = append(resp, connStruct)
	}

	response2.SuccessWithData(resp)
}

// 注销已登录的连接
func DeleteConnectionByKey(c *gin.Context) {
	var req request2.KeyReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	conn, err := hub.get(req.Key)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	conn.close()
	response2.Success()
}