package redis

import (
	"anew-server/pkg/common"
	"fmt"
)

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func (this *InterfaceResult) Unwrap() interface{} {
	if this.Err != nil {
		common.Log.Error(fmt.Sprintf("获取缓存错误: %s"), this.Err)
	}
	return this.Result
}

//获取key对应的值，否则返回指定值
func (this *InterfaceResult) Unwrap_Or(v interface{}) interface{} {
	if this.Err != nil {
		return v
	}
	return this.Result
}