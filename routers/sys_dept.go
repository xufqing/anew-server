package routers

import (
	"anew-server/api/v1/system"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 部门路由
func InitDeptRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", system.GetDepts)
		router.POST("/create", system.CreateDept)
		router.PATCH("/update/:deptId", system.UpdateDeptById)
		router.DELETE("/delete", system.BatchDeleteDeptByIds)
	}
	return router
}
