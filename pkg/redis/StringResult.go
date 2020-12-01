package redis

import (
	"anew-server/pkg/common"
	"fmt"
)

type StringResult struct {
	Result string
	Err    error
}

// 获取字符串结果
func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

// 获取key对应的值
func (this *StringResult) Unwrap() string {
	if this.Err != nil {
		common.Log.Debug(fmt.Sprintf("缓存未击中: %s"), this.Err)
	}
	return this.Result
}

//获取key对应的值，若err则返回指定string
func (this *StringResult) Unwrap_Or(str string) string {
	if this.Err != nil {
		common.Log.Debug(fmt.Sprintf("缓存未击中: %s,返回默认值: %s"), this.Err, str)
		return str
	}
	return this.Result
}

//获取key对应的值，若err则返回指定func
func (this *StringResult) Unwrap_Or_Else(f func() string, key string) string {
	if this.Err != nil {
		//common.Log.Debug(fmt.Sprintf("缓存未击中: %s", key))
		return f()
	}
	//common.Log.Debug(fmt.Sprintf("缓存击中: %s", key))
	return this.Result
}
