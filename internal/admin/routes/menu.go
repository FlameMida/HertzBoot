package routes

import (
	"HertzBoot/internal/admin/http/controllers/admin/v1"
	"HertzBoot/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitMenuRouter(Router *route.RouterGroup) {
	r := Router.Group("menu").Use(middleware.Operations())
	record := Router.Group("menu")
	var menuApi = new(v1.AuthorityMenu)
	{
		r.POST("addBaseMenu", menuApi.AddBaseMenu)           // 新增菜单
		r.POST("addMenuAuthority", menuApi.AddMenuAuthority) //	增加menu和角色关联关系
		r.POST("deleteBaseMenu", menuApi.DeleteBaseMenu)     // 删除菜单
		r.POST("updateBaseMenu", menuApi.UpdateBaseMenu)     // 更新菜单
	}
	{
		record.POST("getMenu", menuApi.GetMenu)                   // 获取菜单树
		record.POST("getMenuList", menuApi.GetMenuList)           // 分页获取基础menu列表
		record.POST("getBaseMenuTree", menuApi.GetBaseMenuTree)   // 获取用户动态路由
		record.POST("getMenuAuthority", menuApi.GetMenuAuthority) // 获取指定角色menu
		record.POST("getBaseMenuById", menuApi.GetBaseMenuById)   // 根据id获取菜单
	}
}
