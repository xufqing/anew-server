package asset

import (
	"anew-server/api/request"
	"anew-server/api/response"
	"anew-server/dao"
	"anew-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取SSH记录列表
func GetSSHRecords(c *gin.Context) {
	// 绑定参数
	var req request.SSHRecordReq
	reqErr := c.Bind(&req)
	if reqErr != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	hostId := utils.Str2Uint(c.Param("hostId"))
	if hostId == 0 {
		response.FailWithMsg("主机编号不正确")
		return
	}
	var err error
	// 创建服务
	s := dao.New()
	records, err := s.GetSSHRecords(hostId,&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var respStruct []response.SSHRecordListResp
	utils.Struct2StructByJson(records, &respStruct)
	// 返回分页数据
	var resp response.PageData
	// 设置分页参数
	resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.DataList = respStruct
	response.SuccessWithData(resp)
}

// 批量删除操作日志
func BatchDeleteSSHRecordByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	// 创建服务
	s := dao.New()
	// 删除数据
	err = s.DeleteSSHRecordByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

func PlaySSHRecord(c *gin.Context) {
	connect_id := c.Query("record")
	s := dao.New()
	record, err := s.GetSSHRecordByConnectID(connect_id)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.cast", connect_id))
	c.Writer.Header().Set("Content-Type", "application/x-asciicast")
	c.String(200, utils.Bytes2Str(utils.ZlibUnCompress(record.Records)))
}
