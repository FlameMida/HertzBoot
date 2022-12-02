package routes

import (
	"HertzBoot/internal/api/http/controllers/admin/v1"
	"HertzBoot/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

func AdminApiRouter(Router *route.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.Operations())
	apiRouterWithoutRecord := Router.Group("api")
	var API = new(v1.Api)
	{
		apiRouter.POST("createApi", API.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", API.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", API.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", API.UpdateApi)               // 更新api
		apiRouter.DELETE("deleteApisByIds", API.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", API.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.POST("getApiList", API.GetApiList) // 获取Api列表
	}
}
