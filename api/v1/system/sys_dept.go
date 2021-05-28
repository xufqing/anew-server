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

// 查询所有部门
func GetDepts(c *gin.Context) {
	// 绑定参数
	var req request2.DeptListReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	// 创建服务
	s := service2.New()
	depts := s.GetDepts(&req)
	if (req.Name != "" || req.Status != nil){
		var newResp []response2.DictTreeResp
		utils.Struct2StructByJson(depts, &newResp)
		response2.SuccessWithData(newResp)
	} else {
		var resp []response2.DeptTreeResp
		resp = service2.GenDeptTree(nil,depts)
		response2.SuccessWithData(resp)
	}
}

// 创建部门
func CreateDept(c *gin.Context) {
	user := GetCurrentUserFromCache(c)
	// 绑定参数
	var req request2.CreateDeptReq
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
	err = s.CreateDept(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 更新部门
func UpdateDeptById(c *gin.Context) {
	// 绑定参数
	var req request2.UpdateDeptReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	deptId := utils.Str2Uint(c.Param("deptId"))
	if deptId == 0 {
		response2.FailWithMsg("部门编号不正确")
		return
	}
	// 创建服务
	s := service2.New()
	// 更新数据
	err = s.UpdateDeptById(deptId, req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 批量删除部门
func BatchDeleteDeptByIds(c *gin.Context) {
	var req request2.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	// 创建服务
	s := service2.New()
	// 删除数据
	err = s.DeleteDeptByIds(req.Ids)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}
