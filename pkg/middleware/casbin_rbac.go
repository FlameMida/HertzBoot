package middleware

import (
	"HertzBoot/internal/core/http/requests"
	"HertzBoot/internal/core/service"
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/response"

	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

var casbinService = new(service.CasbinService)

// CasbinHandler 拦截器
func CasbinHandler() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		claims, _ := ctx.Get("claims")
		waitUse := claims.(*requests.CustomClaims)
		// 获取请求的URI
		obj := ctx.Request.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.CONFIG.System.Env == "develop" || success || !global.CONFIG.Casbin.ApiLevel {
			ctx.Next(c)
		} else {
			response.FailWithDetailed(utils.H{}, "权限不足", ctx)
			ctx.Abort()
			return
		}
	}
}
