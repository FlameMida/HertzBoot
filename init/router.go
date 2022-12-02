package init

import (
	_ "HertzBoot/docs"
	admin "HertzBoot/internal/admin/routes"
	api "HertzBoot/internal/api/routes"
	core "HertzBoot/internal/core/routes"
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/middleware"
	"context"
	"github.com/FlameMida/accessLog"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzSwagger "github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"time"
)

// 初始化总路由

func Routers(address string) *server.Hertz {
	var Router = server.Default(
		server.WithHostPorts(address),
		server.WithReadTimeout(time.Minute*2),
	)
	Router.Use(accessLog.Logger())
	// 静态资源处理
	Router.Static(global.CONFIG.Local.Path, global.CONFIG.Local.Path)
	// 跨域
	Router.Use(middleware.Cors())
	hlog.Info("use middleware cors")

	Router.GET("/swagger/*any", hertzSwagger.WrapHandler(swaggerFiles.Handler))
	hlog.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	AdminGroup := Router.Group("admin-api")
	{
		// 健康监测
		AdminGroup.GET("/health", func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(200, "ok")
		})

		core.AdminBaseRouter(AdminGroup)
		core.AdminInitDBRouter(AdminGroup)

		//挂载jwt鉴权 挂载RBAC
		AdminGroup.Use(middleware.JWTAuth(), middleware.CasbinHandler())

		core.AdminCasbinRouter(AdminGroup)
		core.AdminOperationsRouter(AdminGroup)
		core.AdminSystemRouter(AdminGroup)

		admin.InitAdminRouter(AdminGroup)
		admin.InitAuthorityRouter(AdminGroup)
		admin.InitMenuRouter(AdminGroup)

		api.AdminApiRouter(AdminGroup)

	}

	hlog.Info("router registration success")
	return Router
}
