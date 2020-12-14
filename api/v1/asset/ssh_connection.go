package asset

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"github.com/gin-gonic/gin"
)

// 获取连接列表
func GetConnections(c *gin.Context) {
	var resp []response.ConnectionResp
	for client, _ := range hub.Clients {
		var connStruct response.ConnectionResp
		connStruct.Key = hub.Clients[client].Key
		connStruct.UserName = hub.Clients[client].UserName
		connStruct.Name = hub.Clients[client].Name
		connStruct.HostName = hub.Clients[client].HostName
		connStruct.IpAddress = hub.Clients[client].IpAddress
		connStruct.Port = hub.Clients[client].Port
		connStruct.ConnectTime = hub.Clients[client].ConnectTime
		connStruct.LastActiveTime = hub.Clients[client].LastActiveTime
		resp = append(resp, connStruct)
	}

	response.SuccessWithData(resp)
}

// 注销已登录的连接
func DeleteConnectionByKey(c *gin.Context) {
	var req request.KeyReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	conn, err := hub.get(req.Key)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 连接池删除对象
	hub.delete(req.Key)
	// 关闭socket
	conn.Conn.Close()
	response.Success()
}