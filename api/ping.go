package api

import (
	"github.com/gin-gonic/gin"
)

// 检查服务器是否通畅
func Ping(c *gin.Context) {
	c.JSON(200,gin.H{
		"msg":"ok",
	})
}

