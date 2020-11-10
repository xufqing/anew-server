package routers

import (
	"anew-server/api/v1/system"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 菜单路由
func InitOperLogRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("operlog").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", system.GetOperLogs)
		router.DELETE("/delete", system.BatchDeleteOperLogByIds)
	}
	return router
}
