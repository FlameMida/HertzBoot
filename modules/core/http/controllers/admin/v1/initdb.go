package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/response"
	"HertzBoot/modules/core/http/requests"
	"HertzBoot/modules/core/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"
)

type DB struct {
}

var initDBService = new(service.InitDBService)

// InitDB
// @Tags    Admin.InitDB
// @Summary 初始化用户数据库
// @Produce application/json
// @Param   data body     requests.InitDB true "初始化数据库参数"
// @Success 200  {string} string          "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router  /admin-api/init/initDB [post]
func (i *DB) InitDB(c context.Context, ctx *app.RequestContext) {
	if global.DB != nil {
		global.LOG.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", ctx)
		return
	}
	var dbInfo requests.InitDB
	if err := ctx.BindAndValidate(&dbInfo); err != nil {
		global.LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.LOG.Error("自动创建数据库失败!", zap.Error(err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", ctx)
		return
	}
	response.OkWithData("自动创建数据库成功", ctx)
}

// CheckDB
// @Tags    Admin.InitDB
// @Summary 初始化用户数据库
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router  /admin-api/init/checkDB [post]
func (i *DB) CheckDB(c context.Context, ctx *app.RequestContext) {
	if global.DB != nil {
		global.LOG.Info("数据库无需初始化")
		response.OkWithDetailed(utils.H{"needInit": false}, "数据库无需初始化", ctx)
		return
	} else {
		global.LOG.Info("前往初始化数据库")
		response.OkWithDetailed(utils.H{"needInit": true}, "前往初始化数据库", ctx)
		return
	}
}
