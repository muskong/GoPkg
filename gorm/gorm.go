package gorm

import (
	"database/sql"
	"sync"
	"time"

	"github.com/muskong/GoService/pkg/zaplog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbOnce     sync.Once
	gormDb     *gorm.DB
	gormConfig *GormConfig
)

type (
	GormConfig struct {
		Driver          string
		Dsn             string
		MaxOpenConns    int
		MaxIdleConns    int
		MaxConnLifetime time.Duration
	}
)

func Init(cfg *GormConfig) {
	dbOnce.Do(cfg.initDbConnection)
}

func ClientNew() *gorm.DB {
	return gormDb
}

func (g *GormConfig) initDbConnection() {
	var err error
	sqlDB, err := sql.Open(g.Driver, g.Dsn)
	if err != nil {
		zaplog.Sugar.DPanic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(g.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(g.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(g.MaxConnLifetime)

	gormDb, err = gorm.Open(mysql.New(mysql.Config{
		// DSN: g.Dsn, // data source name
		Conn:                      sqlDB,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		zaplog.Sugar.DPanic(err)
	}

	if gormDb.Error != nil {
		zaplog.Sugar.DPanic(gormDb.Error)
	}

	return
}
