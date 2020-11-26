package initialize

import (
	"anew-server/models/system"
	"anew-server/pkg/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化mysql数据库
func Mysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s&charset=%s&collation=%s",
		common.Conf.Mysql.Username,
		common.Conf.Mysql.Password,
		common.Conf.Mysql.Host,
		common.Conf.Mysql.Port,
		common.Conf.Mysql.Database,
		common.Conf.Mysql.Query,
		common.Conf.Mysql.Charset,
		common.Conf.Mysql.Collation,
	)
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      false,       // 禁用彩色打印
	//	},
	//)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		common.Log.Error(fmt.Sprintf("初始化mysql异常: %v", err))
		panic(fmt.Sprintf("初始化mysql异常: %v", err))
	}
	// 打印所有执行的sql
	//db.LogMode(common.Conf.Mysql.LogMode)
	common.Mysql = db
	// 表结构
	autoMigrate()
	common.Log.Info("Mysql初始化完成")
}

// 自动迁移表结构
func autoMigrate() {
	common.Mysql.AutoMigrate(
		//new(models.SysCasbin),
		new(system.SysUser),
		new(system.SysDept),
		new(system.SysRole),
		new(system.SysMenu),
		new(system.SysApi),
		new(system.SysDict),
		new(system.SysOperLog),
	)
}
