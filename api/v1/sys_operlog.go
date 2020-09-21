package v1

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/utils"
	"github.com/gin-gonic/gin"
)

// 获取操作日志列表
func GetOperationLogs(c *gin.Context) {
	// 绑定参数
	var req request.OperationLogListReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	operationLogs, err := s.GetOperationLogs(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var respStruct []response.OperationLogListResponseStruct
	utils.Struct2StructByJson(operationLogs, &respStruct)
	response.SuccessWithData(respStruct)
}

// 批量删除操作日志
func BatchDeleteOperationLogByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	// 删除数据
	err = s.DeleteOperationLogByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
