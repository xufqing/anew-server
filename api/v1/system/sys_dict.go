package system

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// 查询所有字典
func GetDicts(c *gin.Context) {
	// 绑定参数
	var req request.DictListReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	// 创建服务
	s := service.New()
	dicts := s.GetDicts(&req)
	if req.Key != "" || req.Value != "" || req.Status != nil || req.TypeKey != "" {
		var newResp []response.DictTreeResp
		utils.Struct2StructByJson(dicts, &newResp)
		response.SuccessWithData(newResp)
	} else {
		var resp []response.DictTreeResp
		resp = service.GenDictTree(nil, dicts)
		response.SuccessWithData(resp)
	}
}

// 创建字典
func CreateDict(c *gin.Context) {
	user := GetCurrentUser(c)
	// 绑定参数
	var req request.CreateDictReq
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
	s := service.New()
	err = s.CreateDict(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 更新字典
func UpdateDictById(c *gin.Context) {
	// 绑定参数
	var req request.UpdateDictReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	DictId := utils.Str2Uint(c.Param("DictId"))
	if DictId == 0 {
		response.FailWithMsg("字典编号不正确")
		return
	}
	// 创建服务
	s := service.New()
	// 更新数据
	err = s.UpdateDictById(DictId, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 批量删除字典
func BatchDeleteDictByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New()
	// 删除数据
	err = s.DeleteDictByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
