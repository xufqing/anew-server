package system

import (
	"anew-server/api/request"
	"anew-server/api/response"
	"anew-server/dao"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// 获取接口列表
func GetApis(c *gin.Context) {

	// 创建服务
	s := dao.New()
	apis, err := s.GetApis()
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转换成树结构
	var resp []response.ApiListResp
	// 转换成树结构
	resp = dao.GenApiTree(nil, apis)
	response.SuccessWithData(resp)
}

// 创建接口
func CreateApi(c *gin.Context) {
	user := GetCurrentUserFromCache(c)
	// 绑定参数
	var req request.CreateApiReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 参数校验
	err = common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 记录当前创建人信息
	req.Creator = user.(system.SysUser).Name
	// 创建服务
	s := dao.New()
	err = s.CreateApi(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 更新接口
func UpdateApiById(c *gin.Context) {
	// 绑定参数
	var req request.UpdateApiReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 获取path中的apiId
	apiId := utils.Str2Uint(c.Param("apiId"))
	if apiId == 0 {
		response.FailWithMsg("接口编号不正确")
		return
	}
	// 创建服务
	s := dao.New()
	// 更新数据
	err = s.UpdateApiById(apiId, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 批量删除接口
func BatchDeleteApiByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := dao.New()
	// 删除数据
	err = s.DeleteApiByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
