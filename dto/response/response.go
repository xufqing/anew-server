package response

import (
	"github.com/gin-gonic/gin"
)

// http请求响应封装
type RespInfo struct {
	Code int         `json:"code"` // 错误代码代码
	Data interface{} `json:"data"` // 数据内容
	Msg  string      `json:"msg"`  // 消息提示
}

func Result(code int, data interface{}) {
	// 结果以panic异常的形式抛出, 交由异常处理中间件处理
	panic(RespInfo{
		Code: code,
		Data: data,
		Msg:  CustomError[code],
	})
}
func MsgResult(code int, msg string, data interface{}) {
	// 结果以panic异常的形式抛出, 交由异常处理中间件处理
	panic(RespInfo{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func GetMsgResult(code int, msg string, data interface{}) RespInfo {
	return RespInfo{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func GetResult(code int, data interface{}) RespInfo {
	return RespInfo{
		Code: code,
		Data: data,
		Msg:  CustomError[code],
	}
}

func Success() {
	Result(Ok, map[string]interface{}{})
}

func GetSuccess() RespInfo {
	return GetResult(Ok, map[string]interface{}{})
}

func SuccessWithData(data interface{}) {
	Result(Ok, data)
}

func GetSuccessWithData(data interface{}) RespInfo {
	return GetResult(Ok, data)
}

func FailWithMsg(msg string) {
	MsgResult(NotOk, msg, map[string]interface{}{})
}

func GetFailWithMsg(msg string) RespInfo {
	return GetMsgResult(NotOk, msg,map[string]interface{}{})
}

func FailWithCode(code int) {
	// 查找给定的错误码存在对应的错误信息, 默认使用NotOk
	Result(code, map[string]interface{}{})
}

func GetFailWithCode(code int) RespInfo {

	return GetResult(code, map[string]interface{}{})
}

// 写入json返回值
func JSON(c *gin.Context, code int, resp interface{}) {
	// 调用gin写入json
	c.JSON(code, resp)
	// 保存响应对象到context, Operation Log会读取到
	//c.Set(global.Conf.System.OperationLogKey, resp)
}
