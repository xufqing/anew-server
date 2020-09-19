package v1

import (
	"anew-server/common"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 查询当前用户菜单树
func GetMenuTree(c *gin.Context) {
	_, roleIds := GetCurrentUser(c)
	menuList := make([]models.SysMenu, 0)
	for _, roleId := range roleIds {
		// 创建服务
		s := service.New(c)
		menus, err := s.GetMenuList(roleId)
		if err != nil {
			response.FailWithMsg(err.Error())
			return
		}
		for _, menu := range menus {
			menuList = append(menuList, menu)
		}
	}
	// menuList结构体切片去重
	resultMap := map[string]bool{}
	for _, v := range menuList {
		data, _ := json.Marshal(v)
		resultMap[string(data)] = true
	}
	menusResult := []models.SysMenu{}
	for k := range resultMap {
		var t models.SysMenu
		json.Unmarshal([]byte(k), &t)
		menusResult = append(menusResult, t)
	}
	// 转换成树结构
	menusResult = service.GenMenuTree(nil, menusResult)
	// 转为MenuTreeResponseStruct
	var resp []response.MenuTreeResp
	utils.Struct2StructByJson(menusResult, &resp)
	response.SuccessWithData(resp)
}

// 查询指定角色的菜单树
func GetAllMenuByRoleId(c *gin.Context) {
	// 创建服务
	s := service.New(c)
	menus, ids, err := s.GetAllMenuByRoleId(utils.Str2Uint(c.Param("roleId")))
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	var resp response.MenuTreeWithAccessResp
	resp.AccessIds = ids
	utils.Struct2StructByJson(menus, &resp.List)
	response.SuccessWithData(resp)
}

// 查询所有菜单
func GetMenus(c *gin.Context) {
	// 创建服务
	s := service.New(c)
	menus := s.GetMenus()
	// 转为MenuTreeResponseStruct
	var resp []response.MenuTreeResp
	utils.Struct2StructByJson(menus, &resp)
	response.SuccessWithData(resp)
}

// 创建菜单
func CreateMenu(c *gin.Context) {
	user, _ := GetCurrentUser(c)
	// 绑定参数
	var req request.CreateMenuReq
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
	err = s.CreateMenu(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 更新菜单
func UpdateMenuById(c *gin.Context) {
	// 绑定参数
	var req gin.H
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 获取path中的menuId
	menuId := utils.Str2Uint(c.Param("menuId"))
	if menuId == 0 {
		response.FailWithMsg("菜单编号不正确")
		return
	}
	// 创建服务
	s := service.New(c)
	// 更新数据
	err = s.UpdateMenuById(menuId, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 批量删除菜单
func BatchDeleteMenuByIds(c *gin.Context) {
	var req request.Req
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	// 删除数据
	err = s.DeleteMenuByIds(req.GetUintIds())
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
