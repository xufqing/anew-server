package common

import (
	"github.com/gobuffalo/packr/v2"
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
)

// 全局常量
const (
	// 本地时间格式
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)