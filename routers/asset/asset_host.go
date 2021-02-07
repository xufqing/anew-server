package asset

import (
	"anew-server/api/v1/asset"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	// 创建SSh连接池
	asset.StartConnectionHub()
	router := r.Group("host").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", asset.GetHosts)
		router.GET("/info/:hostId", asset.GetHostInfo)
		router.POST("/create", asset.CreateHost)
		router.PATCH("/update/:hostId", asset.UpdateHostById)
		router.DELETE("/delete", asset.BatchDeleteHostByIds)
		router.GET("/ssh", asset.SShTunnel)
		router.GET("/ssh/ls", asset.GetPathFromSSh)
		router.POST("/ssh/upload", asset.UploadFileToSSh)
		router.GET("/ssh/download", asset.DownloadFileFromSSh)
		router.DELETE("/ssh/rm", asset.DeleteFileInSSh)
		router.GET("/connection/list", asset.GetConnections)
		router.DELETE("/connection/delete", asset.DeleteConnectionByKey)
		router.GET("/record/list", asset.GetSShRecords)
		router.DELETE("/record/delete", asset.BatchDeleteSShRecordByIds)
		router.GET("/record/download", asset.DownloadSShRecord)
		router.GET("/group/list", asset.GetAssetGroups)
		router.POST("/group/create", asset.CreateAssetGroup)
		router.PATCH("/group/update/:groupId", asset.UpdateAssetGroupByID)
		router.DELETE("/group/delete", asset.BatchDeleteAssetGroupByIds)
	}
	return router
}