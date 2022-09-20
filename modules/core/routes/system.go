package routes

import (
	"HertzBoot/app/middleware"
	"HertzBoot/modules/core/http/controllers/admin/v1"
	"github.com/cloudwego/hertz/pkg/route"
)

func AdminSystemRouter(Router *route.RouterGroup) {
	r := Router.Group("system").Use(middleware.Operations())
	var systems = new(v1.Systems)
	{
		r.POST("getSystemConfig", systems.GetSystemConfig) // 获取配置文件内容
		r.POST("setSystemConfig", systems.SetSystemConfig) // 设置配置文件内容
		r.POST("getServerInfo", systems.GetServerInfo)     // 获取服务器信息
	}
}
