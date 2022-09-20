package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/request"
	"HertzBoot/app/response"
	"HertzBoot/modules/admin/entities"
	"HertzBoot/modules/admin/http/requests"
	"HertzBoot/modules/admin/http/responses"
	"HertzBoot/modules/admin/service"
	"HertzBoot/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"go.uber.org/zap"
)

type AuthorityMenu struct {
}

var baseMenuService = new(service.BaseMenuService)

// GetMenu
// @Tags     Admin.AuthorityMenu
// @Summary  获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body     request.Empty true "空"
// @Success  200  {string} string        "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/menu/getMenu [post]
func (a *AuthorityMenu) GetMenu(c context.Context, ctx *app.RequestContext) {
	if err, menus := menuService.GetMenuTree(tools.GetUserAuthorityId(c, ctx)); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		if menus == nil {
			menus = []entities.Menu{}
		}
		response.OkWithDetailed(responses.SysMenusResponse{Menus: menus}, "获取成功", ctx)
	}
}

// GetBaseMenuTree
// @Tags     Admin.AuthorityMenu
// @Summary  获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body     request.Empty true "空"
// @Success  200  {string} string        "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/menu/getBaseMenuTree [post]
func (a *AuthorityMenu) GetBaseMenuTree(c context.Context, ctx *app.RequestContext) {
	if err, menus := menuService.GetBaseMenuTree(); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(responses.SysBaseMenusResponse{Menus: menus}, "获取成功", ctx)
	}
}

// AddMenuAuthority
// @Tags     Admin.AuthorityMenu
// @Summary  增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.AddMenuAuthorityInfo true "角色ID"
// @Success  200  {string} string                        "{"success":true,"data":{},"msg":"添加成功"}"
// @Router   /admin-api/menu/addMenuAuthority [post]
func (a *AuthorityMenu) AddMenuAuthority(c context.Context, ctx *app.RequestContext) {
	var authorityMenu requests.AddMenuAuthorityInfo
	_ = ctx.BindAndValidate(&authorityMenu)
	//todo
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", ctx)
	} else {
		response.OkWithMessage("添加成功", ctx)
	}
}

// GetMenuAuthority
// @Tags     Admin.AuthorityMenu
// @Summary  获取指定角色menu
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.GetAuthorityId true "角色ID"
// @Success  200  {string} string                 "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/menu/getMenuAuthority [post]
func (a *AuthorityMenu) GetMenuAuthority(c context.Context, ctx *app.RequestContext) {
	var param request.GetAuthorityId
	_ = ctx.BindAndValidate(&param)
	//todo
	if err, menus := menuService.GetMenuAuthority(&param); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithDetailed(responses.SysMenusResponse{Menus: menus}, "获取失败", ctx)
	} else {
		response.OkWithDetailed(utils.H{"menus": menus}, "获取成功", ctx)
	}
}

// AddBaseMenu
// @Tags     Admin.Menu
// @Summary  新增菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.BaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success  200  {string} string            "{"success":true,"data":{},"msg":"添加成功"}"
// @Router   /admin-api/menu/addBaseMenu [post]
func (a *AuthorityMenu) AddBaseMenu(c context.Context, ctx *app.RequestContext) {
	var menu entities.BaseMenu
	_ = ctx.BindAndValidate(&menu)
	//todo
	if err := menuService.AddBaseMenu(menu); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))

		response.FailWithMessage("添加失败", ctx)
	} else {
		response.OkWithMessage("添加成功", ctx)
	}
}

// DeleteBaseMenu
// @Tags     Admin.Menu
// @Summary  删除菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.GetById true "菜单id"
// @Success  200  {string} string          "{"success":true,"data":{},"msg":"删除成功"}"
// @Router   /admin-api/menu/deleteBaseMenu [post]
func (a *AuthorityMenu) DeleteBaseMenu(c context.Context, ctx *app.RequestContext) {
	var menu request.GetById
	_ = ctx.BindAndValidate(&menu)
	//todo
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateBaseMenu
// @Tags     Admin.Menu
// @Summary  更新菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.BaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success  200  {string} string            "{"success":true,"data":{},"msg":"更新成功"}"
// @Router   /admin-api/menu/updateBaseMenu [post]
func (a *AuthorityMenu) UpdateBaseMenu(c context.Context, ctx *app.RequestContext) {
	var menu entities.BaseMenu
	_ = ctx.BindAndValidate(&menu)
	//todo
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetBaseMenuById
// @Tags     Admin.Menu
// @Summary  根据id获取菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.GetById true "菜单id"
// @Success  200  {string} string          "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/menu/getBaseMenuById [post]
func (a *AuthorityMenu) GetBaseMenuById(c context.Context, ctx *app.RequestContext) {
	var idInfo request.GetById
	_ = ctx.BindAndValidate(&idInfo)
	//todo
	if err, menu := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(responses.SysBaseMenuResponse{Menu: menu}, "获取成功", ctx)
	}
}

// GetMenuList
// @Tags     Admin.Menu
// @Summary  分页获取基础menu列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.PageInfo true "页码, 每页大小"
// @Success  200  {string} string           "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/menu/getMenuList [post]
func (a *AuthorityMenu) GetMenuList(c context.Context, ctx *app.RequestContext) {
	var pageInfo request.PageInfo
	_ = ctx.BindAndValidate(&pageInfo)
	//todo
	if err, menuList, total := menuService.GetInfoList(); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
