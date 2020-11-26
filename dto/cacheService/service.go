package cacheService

import (
	"anew-server/pkg/common"
	"anew-server/pkg/redis"
	"anew-server/pkg/utils"
	"bytes"
	"encoding/gob"
	"time"
)

const (
	SERILIZER_JSON = "json"
	SERILIZER_GOB  = "gob"
)

type DBGettFunc func() interface{}
type RedisService struct {
	Operation *redis.StringOperation // 字符操作类
	Expire    time.Duration          // 过期时间
	DBGetter  DBGettFunc             // 缓存不存在返回的函数
	Serilizer string                 // 序列化方式
}

func New(operation *redis.StringOperation, expire time.Duration, serilizer string) *RedisService {
	return &RedisService{Operation: operation, Expire: expire, Serilizer: serilizer}
}

// 设置缓存
func (this RedisService) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, redis.WithExpire(this.Expire)).Unwrap()
}

// 获取缓存，获取失败则返回func（func为db的逻辑）
func (this RedisService) GetCache(key string) (ret interface{}) {
	obj := this.DBGetter()
	if this.Serilizer == SERILIZER_JSON {
		f := func() string {
			return utils.Struct2Json(obj)
		}
		ret = this.Operation.Get(key).Unwrap_Or_Else(f, key)
		if ret != nil {
			this.SetCache(key, ret) //不存在则set key
		}

	} else if this.Serilizer == SERILIZER_GOB {
		f := func() string {
			var buf = &bytes.Buffer{}
			enc := gob.NewEncoder(buf)
			if err := enc.Encode(obj); err != nil {
				return ""
			}
			return buf.String()
		}
		ret = this.Operation.Get(key).Unwrap_Or_Else(f, key)
		if ret != nil {
			this.SetCache(key, ret) //不存在则set key
		}
	}
	return
}

// 获取缓存并转换为对象
func (this RedisService) GetCacheForObject(key string, obj interface{}) interface{} {
	ret := this.GetCache(key)
	if ret == nil {
		return nil
	}
	if this.Serilizer == SERILIZER_JSON {
		utils.Json2Struct(ret.(string), obj) // Json转结构体
	} else if this.Serilizer == SERILIZER_GOB {
		var buf = &bytes.Buffer{}
		buf.WriteString(ret.(string))
		dec := gob.NewDecoder(buf)
		if dec.Decode(obj) != nil {
			common.Log.Error("gob转换对象失败")
			return nil
		}
	}
	return nil
}
