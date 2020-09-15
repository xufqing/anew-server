package routers

import (
	v1 "anew-server/api/v1"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 菜单路由
func InitMenuRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("menu").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/tree", v1.GetMenuTree)
		router.GET("/all/:roleId", v1.GetAllMenuByRoleId)
		router.GET("/list", v1.GetMenus)
		router.POST("/create", v1.CreateMenu)
		router.PATCH("/update/:menuId", v1.UpdateMenuById)
		router.DELETE("/delete", v1.BatchDeleteMenuByIds)
	}
	return router
}
