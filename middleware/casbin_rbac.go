package middleware

import (
	"anew-server/api/response"
	"anew-server/api/v1/system"
	system2 "anew-server/models/system"
	"anew-server/pkg/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// Casbin中间件, 基于RBAC的权限访问控制模型
func CasbinMiddleware(c *gin.Context) {
	// 获取当前登录用户
	user := system.GetCurrentUserFromCache(c)
	// 当前登录用户的角色关键字作为casbin访问实体sub
	sub := user.(system2.SysUser).Role.Keyword
	// 请求URL路径作为casbin访问资源obj(需先清除path前缀)
	obj := strings.Replace(c.Request.URL.Path, "/"+common.Conf.System.UrlPathPrefix, "", 1)
	// 请求方式作为casbin访问动作act
	act := c.Request.Method
	// 获取casbin策略管理器
	e := common.Casbin
	// 检查策略
	ok, _ := e.Enforce(sub, obj, act)

	if !ok {
		response.FailWithCode(response.Forbidden)
	}
	// 处理请求
	c.Next()
}
