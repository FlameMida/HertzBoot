package tools

import (
	"HertzBoot/app/global"
	"HertzBoot/modules/core/http/requests"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	uuid "github.com/satori/go.uuid"
)

// GetUserID 从Hertz的Context中获取从jwt解析出来的用户ID
func GetUserID(_ context.Context, ctx *app.RequestContext) uint {
	if claims, exists := ctx.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*requests.CustomClaims)
		return waitUse.ID
	}
}

// GetUserUuid 从Hertz的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(_ context.Context, ctx *app.RequestContext) uuid.UUID {
	if claims, exists := ctx.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return uuid.UUID{}
	} else {
		waitUse := claims.(*requests.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Hertz的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(_ context.Context, ctx *app.RequestContext) string {
	if claims, exists := ctx.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*requests.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Hertz的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(_ context.Context, ctx *app.RequestContext) *requests.CustomClaims {
	if claims, exists := ctx.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*requests.CustomClaims)
		return waitUse
	}
}
