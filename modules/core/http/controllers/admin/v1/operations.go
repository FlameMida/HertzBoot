package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/request"
	"HertzBoot/app/response"
	"HertzBoot/modules/core/entities"
	"HertzBoot/modules/core/http/requests"
	"HertzBoot/modules/core/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"
)

type Operations struct {
}

var operationsService = new(service.OperationsService)

// DeleteOperations
//
// @Tags     Admin.Operations
// @Summary  删除操作记录
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Operations true "Operations模型"
// @Success  200  {string} string              "{"success":true,"data":{},"msg":"删除成功"}"
// @Router   /admin-api/operations/deleteOperations [delete]
func (s *Operations) DeleteOperations(c context.Context, ctx *app.RequestContext) {
	var operations entities.Operations
	_ = ctx.BindAndValidate(&operations)
	if err := operationsService.DeleteOperations(operations); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// DeleteOperationsByIds
// @Tags     Admin.Operations
// @Summary  批量删除操作记录
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.IdsReq true "批量删除Operations"
// @Success  200  {string} string         "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router   /admin-api/operations/deleteOperationsByIds [delete]
func (s *Operations) DeleteOperationsByIds(c context.Context, ctx *app.RequestContext) {
	var IDS request.IdsReq
	_ = ctx.BindAndValidate(&IDS)
	if err := operationsService.DeleteOperationsByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", ctx)
	} else {
		response.OkWithMessage("批量删除成功", ctx)
	}
}

// FindOperations
//
// @Tags     Admin.Operations
// @Summary  用id查询操作记录
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Operations true "ID"
// @Success  200  {string} string              "{"success":true,"data":{},"msg":"查询成功"}"
// @Router   /admin-api/operations/findOperations [get]
func (s *Operations) FindOperations(c context.Context, ctx *app.RequestContext) {
	var operations entities.Operations
	_ = ctx.BindAndValidate(&operations)
	//todo
	if err, reOperations := operationsService.GetOperations(operations.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", ctx)
	} else {
		response.OkWithDetailed(utils.H{"reOperations": reOperations}, "查询成功", ctx)
	}
}

// GetOperationsList
//
// @Tags     Admin.Operations
// @Summary  分页获取操作记录列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.OperationsSearch true "页码, 每页大小, 搜索条件"
// @Success  200  {string} string                    "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/operations/getOperationsList [get]
func (s *Operations) GetOperationsList(c context.Context, ctx *app.RequestContext) {
	var pageInfo requests.OperationsSearch
	_ = ctx.BindAndValidate(&pageInfo)
	if err, list, total := operationsService.GetOperationsInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
