package asset

import (
	"anew-server/api/v1/asset"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	// 创建SSH连接池
	router := r.Group("host").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", asset.GetHosts)
		router.GET("/info/:hostId", asset.GetHostInfo)
		router.POST("/create", asset.CreateHost)
		router.PATCH("/update/:hostId", asset.UpdateHostById)
		router.DELETE("/delete", asset.BatchDeleteHostByIds)
		router.GET("/ssh", asset.Connect)
		router.GET("/ssh/ls", asset.GetPathFromSSH)
		router.POST("/ssh/upload", asset.UploadFileToSSH)
		router.GET("/ssh/download", asset.DownloadFileFromSSH)
		router.DELETE("/ssh/rm", asset.DeleteFileInSSH)
		router.GET("/connection/list", asset.GetConnections)
		router.DELETE("/connection/delete", asset.DeleteConnectionByKey)
		router.GET("/record/list/:hostId", asset.GetSSHRecords)
		router.DELETE("/record/delete", asset.BatchDeleteSSHRecordByIds)
		router.GET("/record/play", asset.PlaySSHRecord)
		router.GET("/group/list", asset.GetAssetGroups)
		router.POST("/group/create", asset.CreateAssetGroup)
		router.PATCH("/group/update/:groupId", asset.UpdateAssetGroupByID)
		router.DELETE("/group/delete", asset.BatchDeleteAssetGroupByIds)
	}
	return router
}
