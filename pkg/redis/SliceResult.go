package redis

import (
	"anew-server/pkg/common"
	"fmt"
)

type SliceResult struct {
	Result []interface{}
	Err    error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

// 获取key对应的值
func (this SliceResult) Unwrap() []interface{} {
	if this.Err != nil {
		common.Log.Debug(fmt.Sprintf("缓存未击中: %s"), this.Err)
	}
	return this.Result
}

//获取key对应的值，若err则返回指定切片
func (this SliceResult) Unwrap_Or(v []interface{}) []interface{} {
	if this.Err != nil {
		return v
	}
	return this.Result
}

// key 迭代器
// redis := cache.NewStringOperation()
//	iter := redis.Mget("name","abc","aaw").Iter()
//	for iter.HasNext(){
//		fmt.Println(iter.Next())
//	}
func (this SliceResult) Iter() *Iterator {

	return NewIterator(this.Result)
}
