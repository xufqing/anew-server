package common

import (
	"errors"
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gobuffalo/packr/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

// 全局变量
var (
	// 配置信息
	Conf Configuration
	// packr盒子用于打包配置文件到golang编译后的二进制程序中
	ConfBox *packr.Box
	// zap日志
	Log *zap.SugaredLogger
	// Mysql实例
	Mysql *gorm.DB
	// Redis实例
	Redis *redis.Client
	// validation.v10校验器
	Validate *validator.Validate
	// validation.v10相关翻译器
	Translator ut.Translator
	// Casbin 管理器
	Casbin *casbin.SyncedEnforcer
)

// 全局常量
const (
	// 本地时间格式
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)

// 只返回一个错误即可
func NewValidatorError(err error, custom map[string]string) (e error) {
	if err == nil {
		return
	}
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		tranStr := e.Translate(Translator)
		// 判断错误字段是否在自定义集合中，如果在，则替换错误信息中的字段
		if v, ok := custom[e.Field()]; ok {
			return errors.New(strings.Replace(tranStr, e.Field(), v, 1))
		} else {
			return errors.New(tranStr)
		}
	}
	return
}
