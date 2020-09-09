package initialize

import (
	"anew-server/api"
	"anew-server/common"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(common.Conf.System.AppMode)
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	// r := gin.Default()
	// 创建不带中间件的路由:
	r := gin.New()
	apiGroup := r.Group("api")
	// ping
	apiGroup.GET("/ping", api.Ping)

	// 方便统一添加路由前缀
	return r
}
