package routes

import (
	"HertzBoot/app/middleware"
	"HertzBoot/modules/core/http/controllers/admin/v1"
	"github.com/cloudwego/hertz/pkg/route"
)

func AdminCasbinRouter(Router *route.RouterGroup) {
	r := Router.Group("casbin").Use(middleware.Operations())
	noRecord := Router.Group("casbin")
	var casbinApi = new(v1.Casbin)
	{
		r.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		noRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
