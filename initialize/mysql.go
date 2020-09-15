package initialize

import (
	"anew-server/common"
	"anew-server/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql驱动
	"github.com/jinzhu/gorm"
)

// 初始化mysql数据库
func Mysql() {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s&charset=%s&collation=%s",
		common.Conf.Mysql.Username,
		common.Conf.Mysql.Password,
		common.Conf.Mysql.Host,
		common.Conf.Mysql.Port,
		common.Conf.Mysql.Database,
		common.Conf.Mysql.Query,
		common.Conf.Mysql.Charset,
		common.Conf.Mysql.Collation,
	))
	if err != nil {
		panic(fmt.Sprintf("初始化mysql异常: %v", err))
	}
	// 打印所有执行的sql
	db.LogMode(common.Conf.Mysql.LogMode)
	common.Mysql = db
	// 表结构
	autoMigrate()
	common.Log.Debug("初始化mysql完成")
}

// 自动迁移表结构
func autoMigrate() {
	common.Mysql.AutoMigrate(
		new(models.SysUser),
		new(models.SysRole),
		new(models.SysMenu),
		new(models.SysApi),
		new(models.SysCasbin),
		new(models.SysOperationLog),
	)
}
