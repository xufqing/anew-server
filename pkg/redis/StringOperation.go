package redis

import "time"

// redis.Set("b","xx",redis.WithExpire(time.Second*15),redis.WithNX())
func (this *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPR)
	nx := OperationAttrs(attrs).Find(ATTR_NX).Unwrap_Or(nil)
	if nx != nil {
		// redis锁
		return NewInterfaceResult(this.redis.SetNX(this.ctx, key, value, exp.Unwrap_Or(time.Second*0).(time.Duration)).Result())
	}
	xx := OperationAttrs(attrs).Find(ATTR_XX).Unwrap_Or(nil)
	if xx != nil {
		return NewInterfaceResult(this.redis.SetXX(this.ctx, key, value, exp.Unwrap_Or(time.Second*0).(time.Duration)).Result())
	}
	return NewInterfaceResult(this.redis.Set(this.ctx, key, value, exp.Unwrap_Or(time.Second*0).(time.Duration)).Result())
}

// redis := cache.NewStringOperation()
// fmt.Println(redis.Get("add").Unwrap_Or("aaw"))
func (this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(this.redis.Get(this.ctx, key).Result())

}
func (this *StringOperation) Mget(key ...string) *SliceResult {
	return NewSliceResult(this.redis.MGet(this.ctx, key...).Result())
}
