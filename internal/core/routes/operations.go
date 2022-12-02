package routes

import (
	"HertzBoot/internal/core/http/controllers/admin/v1"
	"HertzBoot/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

func AdminOperationsRouter(Router *route.RouterGroup) {
	r := Router.Group("operations")
	noRecord := Router.Group("operations").Use(middleware.Operations())
	var operations = new(v1.Operations)
	{
		r.DELETE("deleteOperations", operations.DeleteOperations)           // 删除Operations
		r.DELETE("deleteOperationsByIds", operations.DeleteOperationsByIds) // 批量删除Operations
	}
	{
		noRecord.GET("findOperations", operations.FindOperations)       // 根据ID获取Operations
		noRecord.GET("getOperationsList", operations.GetOperationsList) // 获取Operations列表
	}
}
