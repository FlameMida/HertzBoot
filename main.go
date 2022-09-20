package main

import (
	"HertzBoot/app/global"
	"HertzBoot/app/provider"
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
	global.VP = provider.Viper() // 初始化Viper
	global.LOG = provider.Zap()  // 初始化zap日志库
	global.DB = provider.Gorm()  // gorm连接数据库
	provider.Redis()
	provider.Timer()
	if global.DB != nil {
		provider.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
	provider.RunServer()
}
