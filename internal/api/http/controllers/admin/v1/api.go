package v1

import (
	"HertzBoot/internal/api/entities"
	"HertzBoot/internal/api/http/requests"
	"HertzBoot/internal/api/http/responses"
	"HertzBoot/internal/api/service"
	"HertzBoot/pkg/request"
	"HertzBoot/pkg/response"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type Api struct {
}

var apiService = new(service.ApiService)

// CreateApi
// @Tags     Admin.Api
// @Summary  创建基础api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Api true "api路径, api中文描述, api组, 方法"
// @Success  200  {string} string       "{"success":true,"data":{},"msg":"创建成功"}"
// @Router   /admin-api/api/createApi [post]
func (s *Api) CreateApi(c context.Context, ctx *app.RequestContext) {
	var api entities.Api
	if err := ctx.BindAndValidate(&api); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := apiService.CreateApi(api); err != nil {
		hlog.Error("创建失败!", err.Error())
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithMessage("创建成功", ctx)
	}
}

// DeleteApi
// @Tags     Admin.Api
// @Summary  删除api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Api true "ID"
// @Success  200  {string} string       "{"success":true,"data":{},"msg":"删除成功"}"
// @Router   /admin-api/api/deleteApi [post]
func (s *Api) DeleteApi(c context.Context, ctx *app.RequestContext) {
	var api entities.Api
	if err := ctx.BindAndValidate(&api); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := apiService.DeleteApi(api); err != nil {
		hlog.Error("删除失败!", err.Error())
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// GetApiList
// @Tags     Admin.Api
// @Summary  分页获取API列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.SearchApiParams true "分页获取API列表"
// @Success  200  {string} string                   "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/api/getApiList [post]
func (s *Api) GetApiList(c context.Context, ctx *app.RequestContext) {
	var pageInfo requests.SearchApiParams
	if err := ctx.BindAndValidate(&pageInfo); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, list, total := apiService.GetAPIInfoList(pageInfo.Api, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		hlog.Error("获取失败!", err.Error())
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

// GetApiById
// @Tags     Admin.Api
// @Summary  根据id获取api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.GetById true "根据id获取api"
// @Success  200  {string} string          "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/api/getApiById [post]
func (s *Api) GetApiById(c context.Context, ctx *app.RequestContext) {
	var idInfo request.GetById
	if err := ctx.BindAndValidate(&idInfo); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err, api := apiService.GetApiById(idInfo.ID)
	if err != nil {
		hlog.Error("获取失败!", err.Error())
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(responses.SysAPIResponse{Api: api}, ctx)
	}
}

// UpdateApi
// @Tags     Admin.Api
// @Summary  创建基础api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Api true "api路径, api中文描述, api组, 方法"
// @Success  200  {string} string       "{"success":true,"data":{},"msg":"修改成功"}"
// @Router   /admin-api/api/updateApi [post]
func (s *Api) UpdateApi(c context.Context, ctx *app.RequestContext) {
	var api entities.Api
	if err := ctx.BindAndValidate(&api); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := apiService.UpdateApi(api); err != nil {
		hlog.Error("修改失败!", err.Error())
		response.FailWithMessage("修改失败", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetAllApis
// @Tags     Admin.Api
// @Summary  获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/api/getAllApis [post]
func (s *Api) GetAllApis(c context.Context, ctx *app.RequestContext) {
	if err, apis := apiService.GetAllApis(); err != nil {
		hlog.Error("获取失败!", err.Error())
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(responses.SysAPIListResponse{Apis: apis}, "获取成功", ctx)
	}
}

// DeleteApisByIds
// @Tags     Admin.Api
// @Summary  删除选中Api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.IdsReq true "ID"
// @Success  200  {string} string         "{"success":true,"data":{},"msg":"删除成功"}"
// @Router   /admin-api/api/deleteApisByIds [delete]
func (s *Api) DeleteApisByIds(c context.Context, ctx *app.RequestContext) {
	var ids request.IdsReq
	if err := ctx.BindAndValidate(&ids); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := apiService.DeleteApisByIds(ids); err != nil {
		hlog.Error("删除失败!", err.Error())
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
