package routes

import (
	"HertzBoot/app/middleware"
	"HertzBoot/modules/admin/http/controllers/admin/v1"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitAuthorityRouter(Router *route.RouterGroup) {
	record := Router.Group("authority").Use(middleware.Operations())
	r := Router.Group("authority")
	var authorityApi = new(v1.Authority)
	{
		record.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		record.POST("deleteAuthority", authorityApi.DeleteAuthority)   // 删除角色
		record.PUT("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		record.POST("copyAuthority", authorityApi.CopyAuthority)       // 更新角色
		record.POST("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
	}
	{
		r.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
