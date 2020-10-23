package middleware

import (
	"anew-server/api/v1/system"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/pkg/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

// RBAC权限中间件
func PermsMiddleware(c *gin.Context) {
	// 获取当前登录用户
	_ ,roleId:= system.GetCurrentUser(c)
	permsList := make([]models.SysApi,0)
	// 创建服务
	s := service.New(c)
	for _,role := range roleId{
		roleData,_ := s.GetPermsByRoleId(role)
		//如果是管理员直接放行
		//if roleData.Keyword == "admin" {
		//	c.Next()
		//	return
		//}
		for _,api := range roleData.Apis{
			permsList = append(permsList,api)
		}
	}
	// permsList去重
	resultMap := map[string]bool{}
	for _, v := range permsList {
		data, _ := json.Marshal(v)
		resultMap[string(data)] = true
	}
	permsResult := []models.SysApi{}
	for k := range resultMap {
		var t models.SysApi
		json.Unmarshal([]byte(k), &t)
		permsResult = append(permsResult, t)
	}

	// 检查权限
	// 请求URL路径
	obj := strings.Replace(c.FullPath(), "/"+common.Conf.System.UrlPathPrefix, "", 1)
	// 请求方式
	act := c.Request.Method
	pass := false
	for _,prms := range permsResult{
		if obj==prms.Path && act == prms.Method {
			pass = true
			break
		}
	}
	if !pass {
		response.FailWithCode(response.Forbidden)
	}
	// 处理请求
	c.Next()
}