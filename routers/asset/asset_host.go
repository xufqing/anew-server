package asset

import (
	"anew-server/api/v1/asset"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	// 创建Ssh连接池
	asset.StartConnectionHub()
	router := r.Group("host").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", asset.GetHosts)
		router.GET("/info/:hostId", asset.GetHostInfo)
		router.POST("/create", asset.CreateHost)
		router.PATCH("/update/:hostId", asset.UpdateHostById)
		router.DELETE("/delete", asset.BatchDeleteHostByIds)
		router.GET("/ssh", asset.SshTunnel)
		router.GET("/ssh/ls", asset.GetPathFromSsh)
		router.POST("/ssh/upload", asset.UploadFileToSsh)
		router.GET("/ssh/download", asset.DownloadFileFromSsh)
		router.DELETE("/ssh/rm", asset.DeleteFileInSsh)
		router.GET("/connection/list", asset.GetConnections)
		router.DELETE("/connection/delete", asset.DeleteConnectionByKey)
		router.GET("/record/list", asset.GetSshRecords)
		router.DELETE("/record/delete", asset.BatchDeleteSshRecordByIds)
		router.GET("/record/download", asset.DownloadSshRecord)
	}
	return router
}