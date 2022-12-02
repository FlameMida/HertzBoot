package v1

import (
	"HertzBoot/internal/core/entities"
	systemRes "HertzBoot/internal/core/http/responses"
	"HertzBoot/internal/core/service"
	"HertzBoot/pkg/response"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
		hlog.Error("获取失败!", err.Error())
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
	if err := ctx.BindAndValidate(&sys); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := configService.SetSystemConfig(sys); err != nil {
		hlog.Error("设置失败!", err.Error())
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
		hlog.Error("获取失败!", err.Error())
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(utils.H{"server": server}, "获取成功", ctx)
	}
}
