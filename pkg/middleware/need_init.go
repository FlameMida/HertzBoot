package middleware

import (
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/response"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// NeedInit 处理跨域请求,支持options访问
func NeedInit() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		if global.DB == nil {
			response.OkWithDetailed(utils.H{
				"needInit": true,
			}, "前往初始化数据库", ctx)
			ctx.Abort()
		} else {
			ctx.Next(c)
		}
		// 处理请求
	}
}
