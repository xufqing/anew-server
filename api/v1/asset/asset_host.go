package asset

import (
	request2 "anew-server/api/request"
	response2 "anew-server/api/response"
	getuser "anew-server/api/v1/system"
	service2 "anew-server/dao"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// 获取列表
func GetHosts(c *gin.Context) {
	// 绑定参数
	var req request2.HostReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	// 创建服务
	s := service2.New()
	hosts, err := s.GetHosts(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var respStruct []response2.HostListResp
	utils.Struct2StructByJson(hosts, &respStruct)
	// 返回分页数据
	var resp response2.PageData
	// 设置分页参数
	resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.DataList = respStruct
	response2.SuccessWithData(resp)
}

// 创建
func CreateHost(c *gin.Context) {
	user := getuser.GetCurrentUserFromCache(c)
	// 绑定参数
	var req request2.CreateHostReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	// 参数校验
	err = common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	// 记录当前创建人信息
	req.Creator = user.(system.SysUser).Name
	// 创建服务
	s := service2.New()
	err = s.CreateHost(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 获取当前主机信息
func GetHostInfo(c *gin.Context) {
	// 绑定参数
	var req gin.H
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	hostId := utils.Str2Uint(c.Param("hostId"))
	if hostId == 0 {
		response2.FailWithMsg("接口编号不正确")
		return
	}
	// 创建服务
	s := service2.New()
	host,err := s.GetHostById(hostId)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var connStruct response2.HostListResp
	utils.Struct2StructByJson(host, &connStruct)
	response2.SuccessWithData(connStruct)
}

// 更新
func UpdateHostById(c *gin.Context) {
	// 绑定参数
	var req gin.H
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	hostId := utils.Str2Uint(c.Param("hostId"))
	if hostId == 0 {
		response2.FailWithMsg("接口编号不正确")
		return
	}
	// 创建服务
	s := service2.New()
	// 更新数据
	err = s.UpdateHostById(hostId, req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 批删除
func BatchDeleteHostByIds(c *gin.Context) {
	var req request2.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	// 创建服务
	s := service2.New()
	// 删除数据
	err = s.DeleteHostByIds(req.Ids)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}
