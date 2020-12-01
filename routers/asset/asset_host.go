package asset

import (
	"anew-server/api/v1/asset"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("host").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", asset.GetHosts)
		router.POST("/create", asset.CreateHost)
		router.PATCH("/update/:hostId", asset.UpdateHostById)
		router.DELETE("/delete", asset.BatchDeleteHostByIds)
	}
	return router
}