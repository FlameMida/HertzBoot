package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/request"
	"HertzBoot/app/response"
	"HertzBoot/modules/admin/entities"
	"HertzBoot/modules/admin/http/requests"
	"HertzBoot/modules/admin/http/responses"
	"HertzBoot/modules/admin/service"
	coreRequests "HertzBoot/modules/core/http/requests"
	coreService "HertzBoot/modules/core/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"

	"go.uber.org/zap"
)

type Authority struct {
}

var authorityService = new(service.AuthorityService)
var menuService = new(service.MenuService)
var casbinService = new(coreService.CasbinService)

// CreateAuthority
// @Tags     Admin.Authority
// @Summary  创建角色
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Authority true "权限id, 权限名, 父角色id"
// @Success  200  {string} string             "{"success":true,"data":{},"msg":"创建成功"}"
// @Router   /admin-api/authority/createAuthority [post]
func (a *Authority) CreateAuthority(c context.Context, ctx *app.RequestContext) {
	var authority entities.Authority
	_ = ctx.BindAndValidate(&authority)
	// todo
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), ctx)
	} else {
		_ = menuService.AddMenuAuthority(requests.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, coreRequests.DefaultCasbin())
		response.OkWithDetailed(responses.SysAuthorityResponse{Authority: authBack}, "创建成功", ctx)
	}
}

// CopyAuthority
// @Tags     Admin.Authority
// @Summary  拷贝角色
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     responses.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success  200  {string} string                             "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router   /admin-api/authority/copyAuthority [post]
func (a *Authority) CopyAuthority(c context.Context, ctx *app.RequestContext) {
	var copyInfo responses.SysAuthorityCopyResponse
	_ = ctx.BindAndValidate(&copyInfo)
	//todo
	if err, authBack := authorityService.CopyAuthority(copyInfo); err != nil {
		global.LOG.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(responses.SysAuthorityResponse{Authority: authBack}, "拷贝成功", ctx)
	}
}

// DeleteAuthority
// @Tags     Admin.Authority
// @Summary  删除角色
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Authority true "删除角色"
// @Success  200  {string} string             "{"success":true,"data":{},"msg":"删除成功"}"
// @Router   /admin-api/authority/deleteAuthority [post]
func (a *Authority) DeleteAuthority(c context.Context, ctx *app.RequestContext) {
	var authority entities.Authority
	_ = ctx.BindAndValidate(&authority)
	//todo
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateAuthority
// @Tags     Admin.Authority
// @Summary  更新角色信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Authority true "权限id, 权限名, 父角色id"
// @Success  200  {string} string             "{"success":true,"data":{},"msg":"更新成功"}"
// @Router   /admin-api/authority/updateAuthority [post]
func (a *Authority) UpdateAuthority(c context.Context, ctx *app.RequestContext) {
	var auth entities.Authority
	_ = ctx.BindAndValidate(&auth)
	//todo
	if err, authority := authorityService.UpdateAuthority(auth); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(responses.SysAuthorityResponse{Authority: authority}, "更新成功", ctx)
	}
}

// GetAuthorityList
// @Tags     Admin.Authority
// @Summary  分页获取角色列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.PageInfo true "页码, 每页大小"
// @Success  200  {string} string           "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/authority/getAuthorityList [post]
func (a *Authority) GetAuthorityList(c context.Context, ctx *app.RequestContext) {
	var pageInfo request.PageInfo
	_ = ctx.BindAndValidate(&pageInfo)
	//todo
	if err, list, total := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// SetDataAuthority
// @Tags     Admin.Authority
// @Summary  设置角色资源权限
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Authority true "设置角色资源权限"
// @Success  200  {string} string             "{"success":true,"data":{},"msg":"设置成功"}"
// @Router   /admin-api/authority/setDataAuthority [post]
func (a *Authority) SetDataAuthority(c context.Context, ctx *app.RequestContext) {
	var auth entities.Authority
	_ = ctx.BindAndValidate(&auth)
	//todo
	if err := authorityService.SetDataAuthority(auth); err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}
