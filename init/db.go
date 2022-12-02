package init

import (
	"HertzBoot/init/internal"
	adminEnt "HertzBoot/internal/admin/entities"
	apiEnt "HertzBoot/internal/api/entities"
	"HertzBoot/internal/core/entities"
	"HertzBoot/pkg/global"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@author: Flame
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// DB
// @author:      Flame
// @function:    DB
// @description: 初始化数据库并产生数据库全局变量
// @return:      *sqlx.DB
func DB() (db *sqlx.DB) {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		panic("database is empty in config file")

	}
	//dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	dsn := m.Dsn()
	// 也可以使用MustConnect连接不成功就panic
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("connect to DB failed: " + err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetMaxIdleConns(m.MaxIdleConns)
	return
}

// MysqlTables
//@author: Flame
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		adminEnt.Admin{},
		adminEnt.Authority{},
		apiEnt.Api{},
		adminEnt.BaseMenu{},
		adminEnt.BaseMenuParameter{},
		entities.JwtBlacklist{},
		entities.Operations{},
	)
	if err != nil {
		hlog.Error("register table failed", err.Error())
		os.Exit(0)
	}
	hlog.Info("register table success")
}

//@author: Flame
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		hlog.Error("MySQL启动异常", err.Error())
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: Flame
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		config.Logger = internal.Default.LogMode(logger.Info)
	}
	return config
}
