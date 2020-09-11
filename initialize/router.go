package initialize

import (
	"anew-server/common"
	"anew-server/routers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(common.Conf.System.AppMode)
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	// r := gin.Default()
	// 创建不带中间件的路由:
	r := gin.New()
	apiGroup := r.Group(common.Conf.System.UrlPathPrefix)
	// 注册公共路由，所有人都可以访问
	routers.InitPublicRouter(apiGroup)
	// 方便统一添加路由前缀
	v1 := apiGroup.Group("v1")
	{
		routers.InitUserRouter(v1)
	}


	return r
}
