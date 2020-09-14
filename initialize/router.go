package initialize

import (
	"anew-server/common"
	"anew-server/middleware"
	"anew-server/routers"
	"github.com/gin-gonic/gin"
	"fmt"
)

func Routers() *gin.Engine {
	gin.SetMode(common.Conf.System.AppMode)
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	// r := gin.Default()
	// 创建不带中间件的路由:
	r := gin.New()
	// 添加全局异常处理中间件
	r.Use(middleware.Exception)
	// 添加全局事务处理中间件
	r.Use(middleware.Transaction)
	// 初始化jwt auth中间件
	authMiddleware, err := middleware.InitAuth()

	if err != nil {
		panic(fmt.Sprintf("初始化jwt auth中间件失败: %v", err))
	}
	common.Log.Debug("初始化jwt auth中间件完成")
	apiGroup := r.Group(common.Conf.System.UrlPathPrefix)
	// 注册公共路由，所有人都可以访问
	routers.InitPublicRouter(apiGroup)
	routers.InitAuthRouter(apiGroup, authMiddleware)      // 注册认证路由, 不会鉴权
	// 方便统一添加路由前缀
	v1 := apiGroup.Group("v1")
	{
		routers.InitUserRouter(v1,authMiddleware)   // 注册用户路由
	}


	return r
}
