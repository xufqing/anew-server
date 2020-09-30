package system

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
func GetUserMenuTree(c *gin.Context) {
	_, roleIds := GetCurrentUser(c)

	menuList := make([]models.SysMenu, 0)
	for _, roleId := range roleIds {
		// 创建服务
		s := service.New(c)
		menus, err := s.GetUserMenuList(roleId)
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
	var resp []response.MenuTreeResp
	// 转换成树结构
	resp = service.GenMenuTree(nil, menusResult)
	response.SuccessWithData(resp)
}

// 查询所有菜单
func GetMenus(c *gin.Context) {
	// 创建服务
	s := service.New(c)
	menus := s.GetMenus()
	var resp []response.MenuTreeResp
	resp = service.GenMenuTree(nil,menus)
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
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	// 删除数据
	err = s.DeleteMenuByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
