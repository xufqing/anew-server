package response

import "github.com/gin-gonic/gin"

func Resp(c *gin.Context, code int, resp interface{}) {
	c.JSON(SUCCSE, gin.H{
		"status": code,
		"data":   resp,
		"msg":    Custom_Msg[code],
	})
}