package routers

import (
	"anew-server/api"
	"github.com/gin-gonic/gin"
)

// 公共路由, 任何人可访问
func InitPublicRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	// 测试api可用，ping
	r.GET("/ping", api.Ping)
	r.GET("/avatar", api.ShowUserAvatar)
	return r
}
