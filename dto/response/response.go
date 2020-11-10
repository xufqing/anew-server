package response

import (
	"anew-server/pkg/common"
	"github.com/gin-gonic/gin"
)

// http请求响应封装
type RespInfo struct {
	Code    int         `json:"code"`    // 错误代码代码
	Status  bool        `json:"status"`  // 状态
	Data    interface{} `json:"data"`    // 数据内容
	Message string      `json:"message"` // 消息提示
}
type RespPageInfo struct {
	Code    int         `json:"code"`    // 错误代码代码
	Status  bool        `json:"status"`  // 状态
	Data    interface{} `json:"data"`    // 数据内容
	Message string      `json:"message"` // 消息提示
}

// 分页封装
type PageInfo struct {
	Current  uint  `json:"current" form:"current"`   // 当前页码
	PageSize uint  `json:"pageSize" form:"pageSize"` // 每页显示条数
	Total    int64 `json:"total"`                    // 数据总条数
	All      bool  `json:"all" form:"all"`           // 不使用分页
}

// 带分页数据封装
type PageData struct {
	PageInfo
	DataList interface{} `json:"data"` // 数据列表
}

// 计算limit/offset, 如果需要用到返回的PageSize, PageNum, 务必保证Total值有效
func (s *PageInfo) GetLimit() (int, int) {
	var pageSize int64
	var current int64
	// 传入参数可能不合法, 设置默认值
	// 每页显示条数不能小于1
	if s.PageSize < 1 {
		pageSize = 10
	}
	// 页码不能小于1
	if s.Current < 1 {
		s.Current = 1
	}

	// 如果偏移量比总条数还多
	if s.Total > 0 {
		if pageSize > s.Total {
			pageSize = s.Total
		}
		if current > s.Total {
			current = s.Total
		}
	}

	// 计算最大页码
	maxPageNum := s.Total/pageSize + 1
	if s.Total%pageSize == 0 {
		maxPageNum = s.Total / pageSize
	}
	// 页码不能小于1
	if maxPageNum < 1 {
		maxPageNum = 1
	}

	// 超出最后一页
	if current > maxPageNum {
		current = maxPageNum
	}
	limit := pageSize
	offset := limit * (current - 1)
	return int(limit), int(offset)
}

func Result(code int, status bool, data interface{}) {
	// 结果以panic异常的形式抛出, 交由异常处理中间件处理
	panic(RespInfo{
		Code:    code,
		Status:  status,
		Data:    data,
		Message: CustomError[code],
	})
}

func MsgResult(code int, status bool, msg string, data interface{}) {
	// 结果以panic异常的形式抛出, 交由异常处理中间件处理
	panic(RespInfo{
		Code:    code,
		Status:  status,
		Data:    data,
		Message: msg,
	})
}

func Success() {
	Result(Ok, true, map[string]interface{}{})
}

func SuccessWithData(data interface{}) {
	Result(Ok, true, data)
}

func SuccessWithPageData(data interface{}) {
	Result(Ok, true, data)
}

func SuccessWithMsg(msg string) {
	MsgResult(Ok, true, msg, map[string]interface{}{})
}

func SuccessWithCode(code int) {
	Result(code, true, map[string]interface{}{})
}

func FailWithMsg(msg string) {
	MsgResult(NotOk, false, msg, map[string]interface{}{})
}

func FailWithCode(code int) {
	Result(code, false, map[string]interface{}{})
}

// 写入json返回值
func JSON(c *gin.Context, code int, resp interface{}) {
	// 调用gin写入json
	c.JSON(code, resp)
	// 保存响应对象到context, Operation Log会读取到
	c.Set(common.Conf.System.OperationLogKey, resp)
}
