package routers

import (
	v1 "anew-server/api/v1"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 接口路由
func InitApiRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("api").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", v1.GetApis)
		router.POST("/create", v1.CreateApi)
		router.PATCH("/update/:apiId", v1.UpdateApiById)
		router.DELETE("/delete", v1.BatchDeleteApiByIds)
	}
	return router
}
