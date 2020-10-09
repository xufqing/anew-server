package service

import (
	"anew-server/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MysqlService struct {
	tx *gorm.DB // 事务对象实例
	db *gorm.DB // 无事务对象实例
}

// 初始化服务
func New(c *gin.Context) MysqlService {
	// 获取事务对象
	tx := common.GetTx(c)
	return MysqlService{
		tx: tx,
		db: common.Mysql,
	}
}
