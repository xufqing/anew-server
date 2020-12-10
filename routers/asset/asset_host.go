package asset

import (
	"anew-server/api/v1/asset"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	// 创建SSH连接池
	asset.StartConnectionHub()
	router := r.Group("host").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", asset.GetHosts)
		router.GET("/info/:hostId", asset.GetHostInfo)
		router.POST("/create", asset.CreateHost)
		router.PATCH("/update/:hostId", asset.UpdateHostById)
		router.DELETE("/delete", asset.BatchDeleteHostByIds)
		router.GET("/ssh", asset.SSHTunnel)
		router.GET("/ssh/ls", asset.GetPathFromSSH)
		router.POST("/ssh/upload", asset.UploadFileToSSH)
		router.POST("/ssh/download", asset.DownloadFileFromSSH)
		router.GET("/connection/list", asset.GetConnections)
		router.DELETE("/connection/delete", asset.DeleteConnectionByKey)
	}
	return router
}