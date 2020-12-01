package initialize

import (
	"anew-server/middleware"
	"anew-server/pkg/common"
	"anew-server/routers"
	"anew-server/routers/asset"
	"anew-server/routers/system"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(common.Conf.System.AppMode)
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	// r := gin.Default()
	// 创建不带中间件的路由:
	r := gin.New()
	// 添加访问记录
	r.Use(middleware.AccessLog)
	// 添加操作日志
	r.Use(middleware.OperationLog)
	// 添加全局异常处理中间件
	r.Use(middleware.Exception)
	// 添加跨域中间件, 让请求支持跨域-生产勿用
	// r.Use(middleware.Cors())
	// 初始化jwt auth中间件
	authMiddleware, err := middleware.InitAuth()

	if err != nil {
		panic(fmt.Sprintf("初始化jwt auth中间件失败: %v", err))
	}
	common.Log.Debug("初始化jwt auth中间件完成")
	apiGroup := r.Group(common.Conf.System.UrlPathPrefix)
	// 注册公共路由，所有人都可以访问
	routers.InitPublicRouter(apiGroup)
	system.InitAuthRouter(apiGroup, authMiddleware) // 注册认证路由, 不会鉴权
	// 方便统一添加路由前缀
	v1 := apiGroup.Group("v1")
	{
		system.InitUserRouter(v1, authMiddleware)    // 注册用户路由
		system.InitDeptRouter(v1, authMiddleware)    // 注册部门路由
		system.InitMenuRouter(v1, authMiddleware)    // 注册菜单路由
		system.InitRoleRouter(v1, authMiddleware)    // 注册角色路由
		system.InitApiRouter(v1, authMiddleware)     // 注册接口路由
		system.InitDictRouter(v1, authMiddleware)     // 注册字典路由
		system.InitOperLogRouter(v1, authMiddleware) // 注册操作日志路由
		asset.InitHostRouter(v1, authMiddleware) // 主机管理
	}

	return r
}
