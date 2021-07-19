package asset

import (
	"anew-server/api/request"
	"anew-server/api/response"
	"github.com/gin-gonic/gin"
)

// 获取连接列表
func GetConnections(c *gin.Context) {
	var resp []response.ConnectionResp
	for client, _ := range SteamMap.innerMap {
		var connStruct response.ConnectionResp
		connStruct.ConnectID = SteamMap.innerMap[client].Meta.ConnectId
		connStruct.UserName = SteamMap.innerMap[client].Meta.UserName
		connStruct.HostName = SteamMap.innerMap[client].Meta.HostName
		connStruct.ConnectTime = SteamMap.innerMap[client].CreatedAt
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
	connObj, err := SteamMap.Get(req.Key)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	connObj.Terminal.Close()
	SteamMap.Remove(req.Key)
	response.Success()
}
