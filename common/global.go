package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
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
)

// 全局常量
const (
	// 本地时间格式
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)


// 获取事务对象
func GetTx(c *gin.Context) *gorm.DB {
	// 默认使用无事务的mysql
	tx := Mysql
	if c != nil {
		method := c.Request.Method
		if !(method == "OPTIONS" || method == "GET" || !Conf.System.Transaction) {
			// 从context对象中读取事务对象
			txKey, exists := c.Get("tx")
			if exists {
				if item, ok := txKey.(*gorm.DB); ok {
					tx = item
				}
			}
		}
	}
	return tx
}