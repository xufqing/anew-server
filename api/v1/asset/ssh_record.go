package asset

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models/asset"
	"anew-server/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取ssh记录列表
func GetSShRecords(c *gin.Context) {
	// 绑定参数
	var req request.SShRecordReq
	reqErr := c.Bind(&req)
	if reqErr != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	var sshRecord []asset.SShRecord
	var err error
	// 创建服务
	s := service.New()
	sshRecord, err = s.GetSSHRecords(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	//var respStruct []response.OperationLogListResp
	//utils.Struct2StructByJson(operationLogs, &respStruct)
	// 返回分页数据
	var resp response.PageData
	// 设置分页参数
	resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.DataList = sshRecord
	response.SuccessWithData(resp)
}

// 批量删除操作日志
func BatchDeleteSShRecordByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	// 创建服务
	s := service.New()
	// 删除数据
	err = s.DeleteSSHRecordByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

func DownloadSShRecord(c *gin.Context) {
	record := c.Query("record")
	file := common.Conf.SSh.RecordDir + "/" + record
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", record))
	c.Writer.Header().Set("Content-Type", "application/x-asciicast")
	c.File(file)
}
