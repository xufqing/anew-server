package routers

import (
	v1 "anew-server/api/v1"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// User
func InitUserRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("user").Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/list", v1.GetUsers).Use(authMiddleware.MiddlewareFunc())
		router.POST("/create", v1.CreateUser).Use(authMiddleware.MiddlewareFunc())
		router.PATCH("/update/:userId", v1.UpdateUserById).Use(authMiddleware.MiddlewareFunc())
		router.DELETE("/delete", v1.DeleteUserByIds).Use(authMiddleware.MiddlewareFunc())
	}
	return router
}
