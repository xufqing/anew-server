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

// 查询当前用户菜单树
func GetUserMenuTree(c *gin.Context) {
	user := GetCurrentUserFromCache(c)
	s := service2.New()
	menus, err := s.GetUserMenuList(user.(system.SysUser).RoleId)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	var resp []response2.MenuTreeResp
	// 转换成树结构
	resp = service2.GenMenuTree(nil, menus)
	response2.SuccessWithData(resp)
}

// 查询所有菜单
func GetMenus(c *gin.Context) {
	// 创建服务
	s := service2.New()
	menus := s.GetMenus()
	var resp []response2.MenuTreeResp
	resp = service2.GenMenuTree(nil,menus)
	response2.SuccessWithData(resp)
}

// 创建菜单
func CreateMenu(c *gin.Context) {
	user := GetCurrentUserFromCache(c)
	// 绑定参数
	var req request2.CreateMenuReq
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
	err = s.CreateMenu(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 更新菜单
func UpdateMenuById(c *gin.Context) {
	// 绑定参数
	var req request2.UpdateMenuReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	// 获取path中的menuId
	menuId := utils.Str2Uint(c.Param("menuId"))
	if menuId == 0 {
		response2.FailWithMsg("菜单编号不正确")
		return
	}
	// 创建服务
	s := service2.New()
	// 更新数据
	err = s.UpdateMenuById(menuId, req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 批量删除菜单
func BatchDeleteMenuByIds(c *gin.Context) {
	var req request2.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	// 创建服务
	s := service2.New()
	// 删除数据
	err = s.DeleteMenuByIds(req.Ids)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}
