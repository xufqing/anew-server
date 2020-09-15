package routers

import (
	v1 "anew-server/api/v1"
	"anew-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// User
func InitUserRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.POST("/info", v1.GetUserInfo)
		router.GET("/list", v1.GetUsers)
		router.POST("/create", v1.CreateUser)
		router.PATCH("/update/:userId", v1.UpdateUserById)
		router.PUT("/changePwd", v1.ChangePwd)
		router.DELETE("/delete", v1.DeleteUserByIds)
	}
	return router
}
