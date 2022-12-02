package v1

import (
	"HertzBoot/internal/admin/entities"
	"HertzBoot/internal/admin/http/requests"
	"HertzBoot/internal/admin/http/responses"
	"HertzBoot/internal/admin/service"
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/request"
	"HertzBoot/pkg/response"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

type Admin struct {
}

var userService = new(service.AdminService)

// Register
// @Tags    Admin.Admin
// @Summary 用户注册账号
// @Produce application/json
// @Param   data body     requests.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200  {string} string            "{"success":true,"data":{},"msg":"注册成功"}"
// @Router  /admin-api/admin/register [post]
func (a *Admin) Register(c context.Context, ctx *app.RequestContext) {
	var r requests.Register
	if err := ctx.BindAndValidate(&r); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	var authorities []entities.Authority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, entities.Authority{
			AuthorityId: v,
		})
	}
	admin := &entities.Admin{
		Username:    r.Username,
		NickName:    r.NickName,
		Password:    r.Password,
		Avatar:      r.Avatar,
		AuthorityId: r.AuthorityId,
		Authorities: authorities,
	}
	err, userReturn := userService.Register(*admin)
	if err != nil {
		hlog.Error("注册失败!", err.Error())
		response.FailWithDetailed(responses.UserResponse{User: userReturn}, "注册失败", ctx)
	} else {
		response.OkWithDetailed(responses.UserResponse{User: userReturn}, "注册成功", ctx)
	}
}

// ChangePassword
// @Tags     Admin.Admin
// @Summary  用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body     requests.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success  200  {string} string                        "{"success":true,"data":{},"msg":"修改成功"}"
// @Router   /admin-api/admin/changePassword [put]
func (a *Admin) ChangePassword(c context.Context, ctx *app.RequestContext) {
	var body requests.ChangePasswordStruct
	if err := ctx.BindAndValidate(&body); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	admin := &entities.Admin{
		Username: body.Username,
		Password: body.Password,
	}
	if err, _ := userService.ChangePassword(admin, body.NewPassword); err != nil {
		hlog.Error("修改失败!", err.Error())
		response.FailWithMessage("修改失败，原密码与当前账户不符", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetUserList
// @Tags     Admin.Admin
// @Summary  分页获取用户列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.PageInfo true "页码, 每页大小"
// @Success  200  {string} string           "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/admin/getUserList [post]
func (a *Admin) GetUserList(c context.Context, ctx *app.RequestContext) {
	var pageInfo request.PageInfo
	if err := ctx.BindAndValidate(&pageInfo); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
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

// SetUserAuthority
// @Tags     Admin.Admin
// @Summary  更改用户权限
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.SetUserAuth true "用户UUID, 角色ID"
// @Success  200  {string} string               "{"success":true,"data":{},"msg":"修改成功"}"
// @Router   /admin-api/admin/setUserAuthority [post]
func (a *Admin) SetUserAuthority(c context.Context, ctx *app.RequestContext) {
	var sua requests.SetUserAuth
	if err := ctx.BindAndValidate(&sua); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	userID := pkg.GetUserID(c, ctx)
	uuid := pkg.GetUserUuid(c, ctx)
	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		hlog.Error("修改失败!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
	} else {
		claims := pkg.GetUserInfo(c, ctx)
		j := &pkg.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			hlog.Error("修改失败!", err.Error())
			response.FailWithMessage(err.Error(), ctx)
		} else {
			ctx.Header("new-token", token)
			ctx.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
			response.OkWithMessage("修改成功", ctx)
		}

	}
}

// SetUserAuthorities
// @Tags     Admin.Admin
// @Summary  设置用户权限
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     requests.SetUserAuthorities true "用户UUID, 角色ID"
// @Success  200  {string} string                      "{"success":true,"data":{},"msg":"修改成功"}"
// @Router   /admin-api/admin/setUserAuthorities [post]
func (a *Admin) SetUserAuthorities(c context.Context, ctx *app.RequestContext) {
	var sua requests.SetUserAuthorities
	if err := ctx.BindAndValidate(&sua); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		hlog.Error("修改失败!", err.Error())
		response.FailWithMessage("修改失败", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// DeleteUser
// @Tags     Admin.Admin
// @Summary  删除用户
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.GetById true "用户ID"
// @Success  200  {string} string          "{"success":true,"data":{},"msg":"删除成功"}"
// @Router   /admin-api/admin/deleteUser [delete]
func (a *Admin) DeleteUser(c context.Context, ctx *app.RequestContext) {
	var reqId request.GetById
	if err := ctx.BindAndValidate(&reqId); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	jwtId := pkg.GetUserID(c, ctx)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", ctx)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		hlog.Error("删除失败!", err.Error())
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// SetUserInfo
// @Tags     Admin.Admin
// @Summary  设置用户信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     entities.Admin true "ID, 用户名, 昵称, 头像链接"
// @Success  200  {string} string         "{"success":true,"data":{},"msg":"设置成功"}"
// @Router   /admin-api/admin/setUserInfo [put]
func (a *Admin) SetUserInfo(c context.Context, ctx *app.RequestContext) {
	var body entities.Admin
	if err := ctx.BindAndValidate(&body); err != nil {
		hlog.Error("参数校验不通过!", err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, ReqUser := userService.SetUserInfo(body); err != nil {
		hlog.Error("设置失败!", err.Error())
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithDetailed(utils.H{"userInfo": ReqUser}, "设置成功", ctx)
	}
}

// GetUserInfo
// @Tags     Admin.Admin
// @Summary  获取用户信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router   /admin-api/admin/getUserInfo [get]
func (a *Admin) GetUserInfo(c context.Context, ctx *app.RequestContext) {
	uuid := pkg.GetUserUuid(c, ctx)
	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
		hlog.Error("获取失败!", err.Error())
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(utils.H{"userInfo": ReqUser}, "获取成功", ctx)
	}
}
