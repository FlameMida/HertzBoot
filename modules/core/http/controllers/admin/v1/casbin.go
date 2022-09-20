package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/response"
	"HertzBoot/modules/core/http/requests"
	"HertzBoot/modules/core/http/responses"
	"HertzBoot/modules/core/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
)

type Casbin struct {
}

var casbinService = new(service.CasbinService)

// UpdateCasbin
//
// @Tags     Admin.Casbin
// @Summary  更新角色api权限
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.CasbinInReceive true "权限id, 权限模型列表"
// @Success  200  {string} string                   "{"success":true,"data":{},"msg":"更新成功"}"
// @Router   /admin-api/casbin/UpdateCasbin [post]
func (cas *Casbin) UpdateCasbin(c context.Context, ctx *app.RequestContext) {
	var cmr requests.CasbinInReceive
	_ = ctx.BindAndValidate(&cmr)
	//todo
	if err := casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetPolicyPathByAuthorityId
// @Tags     Admin.Casbin
// @Summary  获取权限列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.CasbinInReceive true "权限id, 权限模型列表"
// @Success  200  {string} string                   "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/casbin/getPolicyPathByAuthorityId [post]
func (cas *Casbin) GetPolicyPathByAuthorityId(c context.Context, ctx *app.RequestContext) {
	var casbin requests.CasbinInReceive
	_ = ctx.BindAndValidate(&casbin)
	//todo
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(responses.PolicyPathResponse{Paths: paths}, "获取成功", ctx)
}
