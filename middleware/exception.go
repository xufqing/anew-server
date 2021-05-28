package middleware

import (
	response2 "anew-server/api/response"
	"anew-server/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// 全局异常处理中间件
func Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 判断是否正常http响应结果放通
			if resp, ok := err.(response2.RespInfo); ok {
				response2.JSON(c, response2.Ok, resp)
				c.Abort()
				return
			}
			// 将异常写入日志
			common.Log.Error(fmt.Sprintf("[Exception]未知异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			// 服务器异常
			resp := response2.RespInfo{
				Code:    response2.InternalServerError,
				Data:    map[string]interface{}{},
				Message: response2.CustomError[response2.InternalServerError],
			}
			// 以json方式写入响应
			response2.JSON(c, response2.Ok, resp)
			c.Abort()
			return
		}
	}()
	c.Next()
}
