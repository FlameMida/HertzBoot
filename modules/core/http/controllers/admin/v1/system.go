package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/response"
	"HertzBoot/modules/core/entities"
	systemRes "HertzBoot/modules/core/http/responses"
	"HertzBoot/modules/core/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"go.uber.org/zap"
)

type Systems struct {
}

var configService = new(service.ConfigService)

// GetSystemConfig
// @Tags     Admin.System
// @Summary  获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/system/getSystemConfig [post]
func (s *Systems) GetSystemConfig(c context.Context, ctx *app.RequestContext) {
	if err, config := configService.GetSystemConfig(); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", ctx)
	}
}

// SetSystemConfig
// @Tags     Admin.System
// @Summary  设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body     entities.System true "设置配置文件内容"
// @Success  200  {string} string          "{"success":true,"data":{},"msg":"设置成功"}"
// @Router   /admin-api/system/setSystemConfig [post]
func (s *Systems) SetSystemConfig(c context.Context, ctx *app.RequestContext) {
	var sys entities.System
	_ = ctx.BindAndValidate(&sys)
	if err := configService.SetSystemConfig(sys); err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithData("设置成功", ctx)
	}
}

// GetServerInfo
// @Tags     Admin.System
// @Summary  获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/system/getServerInfo [post]
func (s *Systems) GetServerInfo(c context.Context, ctx *app.RequestContext) {
	if server, err := configService.GetServerInfo(); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(utils.H{"server": server}, "获取成功", ctx)
	}
}
