package system

import (
	request2 "anew-server/api/request"
	response2 "anew-server/api/response"
	service2 "anew-server/dao"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// 查询所有字典
func GetDicts(c *gin.Context) {
	// 绑定参数
	var req request2.DictListReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	// 创建服务
	s := service2.New()
	dicts := s.GetDicts(&req)
	if req.Key != "" || req.Value != "" || req.Status != nil || req.TypeKey != "" {
		var newResp []response2.DictTreeResp
		utils.Struct2StructByJson(dicts, &newResp)
		response2.SuccessWithData(newResp)
	} else {
		var resp []response2.DictTreeResp
		resp = service2.GenDictTree(nil, dicts)
		response2.SuccessWithData(resp)
	}
}

// 创建字典
func CreateDict(c *gin.Context) {
	user := GetCurrentUserFromCache(c)
	// 绑定参数
	var req request2.CreateDictReq
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
	err = s.CreateDict(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 更新字典
func UpdateDictById(c *gin.Context) {
	// 绑定参数
	var req request2.UpdateDictReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	dictId := utils.Str2Uint(c.Param("dictId"))
	if dictId == 0 {
		response2.FailWithMsg("字典编号不正确")
		return
	}
	// 创建服务
	s := service2.New()
	// 更新数据
	err = s.UpdateDictById(dictId, req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 批量删除字典
func BatchDeleteDictByIds(c *gin.Context) {
	var req request2.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	// 创建服务
	s := service2.New()
	// 删除数据
	err = s.DeleteDictByIds(req.Ids)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}
