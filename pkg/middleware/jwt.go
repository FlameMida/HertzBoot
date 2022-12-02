package middleware

import (
	"HertzBoot/internal/core/entities"
	"HertzBoot/internal/core/service"
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/response"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-jwt/jwt/v4"

	myUtils "HertzBoot/pkg"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
	"time"
)

var jwtService = new(service.JwtService)

func JWTAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 我们这里jwt鉴权取头部信息 Authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithDetailed(utils.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(utils.H{"reload": true}, "您的帐户异地登陆或令牌失效", ctx)
			ctx.Abort()
			return
		}
		j := myUtils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == myUtils.TokenExpired {
				response.FailWithDetailed(utils.H{"reload": true}, "授权已过期", ctx)
				ctx.Abort()
				return
			}
			response.FailWithDetailed(utils.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		// if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.Blacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(utils.H{"reload": true}, err.Error(),ctx)
		//	ctx.Abort()
		// }
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Unix(time.Now().Unix()+global.CONFIG.JWT.ExpiresTime, 0))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			if global.CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					hlog.Error("get redis jwt failed", err.Error())
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.Blacklist(entities.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		ctx.Set("claims", claims)
		ctx.Next(c)
	}
}
