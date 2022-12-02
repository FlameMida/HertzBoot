package main

import (
	Initialize "HertzBoot/init"
	"HertzBoot/pkg/global"
	"database/sql"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                      Hertz Swagger API
// @version                    0.0.1
// @description                This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @BasePath                   /
func main() {
	global.VP = Initialize.Viper() // 初始化Viper
	Initialize.Zap()               // 初始化zap日志库
	global.DB = Initialize.Gorm()  // gorm连接数据库
	Initialize.Redis()
	Initialize.Timer()
	if global.DB != nil {
		Initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
	Initialize.RunServer()
}
