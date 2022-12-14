package service

import (
	"HertzBoot/config"
	adminEnt "HertzBoot/internal/admin/entities"
	apiEnt "HertzBoot/internal/api/entities"
	"HertzBoot/internal/core/entities"
	"HertzBoot/internal/core/http/requests"
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
	"HertzBoot/source"
	"database/sql"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @author:      Flame
// @function:    writeConfig
// @description: 回写配置
// @param:       viper *viper.Viper, mysql config.Mysql
// @return:      error

type InitDBService struct {
}

func (initDBService *InitDBService) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.CONFIG.Mysql = mysql
	cs := pkg.StructToMap(global.CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

// @author:      Flame
// @function:    createTable
// @description: 创建数据库(mysql)
// @param:       dsn string, driver string, createSql
// @return:      error

func (initDBService *InitDBService) createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			hlog.Error("关闭db链接错误", err.Error())
			return
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func (initDBService *InitDBService) initDB(InitDBFunctions ...entities.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// @author:      Flame
// @function:    InitDB
// @description: 创建数据库并初始化
// @param:       conf requests.InitDB
// @return:      error

func (initDBService *InitDBService) InitDB(conf requests.InitDB) error {

	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}

	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Dbname:   conf.DBName,
		Username: conf.UserName,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if MysqlConfig.Dbname == "" {
		return nil
	}

	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.DB = db
	}
	tempDB := global.DB
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			adminEnt.Admin{},
			adminEnt.Authority{},
			apiEnt.Api{},
			adminEnt.BaseMenu{},
			adminEnt.BaseMenuParameter{},
			entities.JwtBlacklist{},
			entities.Operations{},
		)
		if err != nil {
			global.DB = nil
			return err
		}
		global.DB = tx
		err = initDBService.initDB(
			source.Admin,
			source.Api,
			source.AuthorityMenu,
			source.Authority,
			source.AuthoritiesMenus,
			source.Casbin,
			source.DataAuthorities,
			source.BaseMenu,
			source.AdminAuthority,
		)
		if err != nil {
			global.DB = nil
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	global.DB = tempDB

	if err = initDBService.writeConfig(global.VP, MysqlConfig); err != nil {
		return err
	}
	return nil
}
