package routers

import (
	v1 "anew-server/api/v1"
	"github.com/gin-gonic/gin"
)

// User
func InitUserRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	r.GET("users", v1.GetUsers)
	r.POST("user/create",v1.CreateUser)
	return r
}
