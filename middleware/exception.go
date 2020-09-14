package middleware

import (
	"anew-server/common"
	"anew-server/dto/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// 全局异常处理中间件
func Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 将异常写入日志
			common.Log.Error(fmt.Sprintf("[Exception]未知异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			// 服务器异常
			resp := response.RespInfo{
				Code: response.InternalServerError,
				Data: map[string]interface{}{},
				Msg:  response.CustomError[response.InternalServerError],
			}
			// 以json方式写入响应
			response.JSON(c, resp.Code, resp)
			c.Abort()
			return
		}
	}()
	c.Next()
}
