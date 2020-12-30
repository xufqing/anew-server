package asset

import (
	"anew-server/dto/cacheService"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models/asset"
	"anew-server/pkg/common"
	"anew-server/pkg/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// 获取ssh记录列表
func GetSshRecords(c *gin.Context) {
	// 绑定参数
	var req request.SshRecordListReq
	reqErr := c.Bind(&req)
	if reqErr != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	var sshRecord []asset.SshRecord
	var err error
	// 创建缓存对象
	cache := cacheService.New(redis.NewStringOperation(), time.Second*1, cacheService.SERILIZER_JSON)
	key := "sshRecord:" + req.Key + ":" + req.UserName + ":" + req.HostName + ":" + req.IpAddress +
		strconv.Itoa(int(req.Current)) + ":" + strconv.Itoa(int(req.PageSize)) + ":" + strconv.Itoa(int(req.Total))

	cache.DBGetter = func() interface{} {
		// 创建服务
		s := service.New()
		sshRecord, err = s.GetSshRecords(&req)
		return sshRecord
	}
	// 获取缓存
	cache.GetCacheForObject(key, &sshRecord)
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
func BatchDeleteSshRecordByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	// 创建服务
	s := service.New()
	// 删除数据
	err = s.DeleteSshRecordByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

func DownloadSshRecord(c *gin.Context) {
	record := c.Query("record")
	file := common.Conf.Ssh.RecordDir + "/" + record
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file))
	c.Writer.Header().Set("Content-Type", "application/x-asciicast")
	c.File(file)
}