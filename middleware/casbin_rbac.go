package middleware

import (
	v1 "anew-server/api/v1"
	"anew-server/common"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"github.com/gin-gonic/gin"
	"strings"
)

// Casbin中间件, 基于RBAC的权限访问控制模型
func CasbinMiddleware(c *gin.Context) {
	// 获取当前登录用户
	user ,_:= v1.GetCurrentUser(c)
	// 当前登录用户的角色关键字作为casbin访问实体sub
	//sub := user.Role.Keyword
	roles := make([]string,0)
	for _,role := range user.Roles{
		roles = append(roles,role.Keyword)
	}
	sub := roles
	// 请求URL路径作为casbin访问资源obj(需先清除path前缀)
	obj := strings.Replace(c.Request.URL.Path, "/"+common.Conf.System.UrlPathPrefix, "", 1)
	// 请求方式作为casbin访问动作act
	act := c.Request.Method
	// 创建服务
	s := service.New(c)
	// 获取casbin策略管理器
	e, err := s.Casbin()
	if err != nil {
		response.FailWithMsg("获取资源访问策略失败")
		return
	}
	// 检查策略
	pass := false
	for _,prms := range sub{
		ok, _ := e.Enforce(prms, obj, act)
		if ok {
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
