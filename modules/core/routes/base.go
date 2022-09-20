package routes

import (
	"HertzBoot/modules/core/http/controllers/admin/v1"
	"github.com/cloudwego/hertz/pkg/route"
)

func AdminBaseRouter(Router *route.RouterGroup) {
	var baseApi = new(v1.Base)
	{
		Router.POST("login", baseApi.AdminLogin)
		Router.POST("captcha", baseApi.Captcha)
		Router.DELETE("logout", baseApi.Logout)
	}
}
