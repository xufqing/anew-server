package dao

import (
	"anew-server/pkg/common"
	"gorm.io/gorm"
)

type MysqlService struct {
	db *gorm.DB // 数据库对象实例
}

// 初始化服务
func New() MysqlService {
	return MysqlService{
		db: common.Mysql,
	}
}
