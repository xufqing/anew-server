package v1

import (
	"anew-server/common"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/utils"
	"github.com/gin-gonic/gin"
)

// 获取接口列表
func GetApis(c *gin.Context) {
	// 绑定参数
	var req request.ApiListReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	apis, err := s.GetApis(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var respStruct []response.ApiListResp
	utils.Struct2StructByJson(apis, &respStruct)
	response.SuccessWithData(respStruct)
}

// 查询指定角色的接口(以分类分组)
func GetAllApiGroupByCategoryByRoleId(c *gin.Context) {
	// 创建服务
	s := service.New(c)
	// 绑定参数
	apis, ids, err := s.GetAllApiGroupByCategoryByRoleId(utils.Str2Uint(c.Param("roleId")))
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	var resp response.ApiTreeWithAccessResponseStruct
	resp.AccessIds = ids
	utils.Struct2StructByJson(apis, &resp.List)
	response.SuccessWithData(resp)
}

// 创建接口
func CreateApi(c *gin.Context) {
	user,_:= GetCurrentUser(c)
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
	req.Creator = user.Name
	// 创建服务
	s := service.New(c)
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
	var req gin.H
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
	s := service.New(c)
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
	s := service.New(c)
	// 删除数据
	err = s.DeleteApiByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
