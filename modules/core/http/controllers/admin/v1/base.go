package v1

import (
	"HertzBoot/app/global"
	"HertzBoot/app/response"

	adminEntities "HertzBoot/modules/admin/entities"
	"HertzBoot/modules/core/entities"
	"HertzBoot/modules/core/http/requests"
	"HertzBoot/modules/core/http/responses"
	"HertzBoot/modules/core/service"
	"HertzBoot/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type Base struct {
}

var loginService = new(service.LoginService)
var jwtService = new(service.JwtService)

// Captcha
// @Tags     Admin.Base
// @Summary  生成验证码
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router   /admin-api/captcha [post]
func (b *Base) Captcha(_ context.Context, ctx *app.RequestContext) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", ctx)
	} else {
		response.OkWithDetailed(responses.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", ctx)
	}
}

// AdminLogin
// @Tags    Admin.Base
// @Summary 后台用户登录
// @Produce application/json
// @Param   data body     requests.AdminLogin true "用户名, 密码, 验证码"
// @Success 200  {string} string              "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router  /admin-api/login [post]
func (b *Base) AdminLogin(c context.Context, ctx *app.RequestContext) {
	var l requests.AdminLogin
	_ = ctx.BindAndValidate(&l)
	//todo
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &adminEntities.Admin{Username: l.Username, Password: l.Password}
		if err, user := loginService.Login(u); err != nil {
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", ctx)
		} else {
			b.tokenNext(c, ctx, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", ctx)
	}
}

// 登录以后签发jwt
func (b *Base) tokenNext(_ context.Context, ctx *app.RequestContext, admin adminEntities.Admin) {
	j := &tools.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := requests.CustomClaims{
		UUID:        admin.UUID,
		ID:          admin.ID,
		NickName:    admin.NickName,
		Username:    admin.Username,
		AuthorityId: admin.AuthorityId,
		BufferTime:  global.CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-1000, 0)),                          // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+global.CONFIG.JWT.ExpiresTime, 0)), // 过期时间 7天  配置文件
			Issuer:    "issuer",                                                                          // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", ctx)
		return
	}
	if !global.CONFIG.System.UseMultipoint {
		response.OkWithDetailed(responses.LoginResponse{
			User:      admin,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", ctx)
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(admin.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, admin.Username); err != nil {
			global.LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", ctx)
			return
		}
		response.OkWithDetailed(responses.LoginResponse{
			User:      admin,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", ctx)
	} else if err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", ctx)
	} else {
		var blackJWT entities.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.Blacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", ctx)
			return
		}
		if err := jwtService.SetRedisJWT(token, admin.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", ctx)
			return
		}
		response.OkWithDetailed(responses.LoginResponse{
			User:      admin,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", ctx)
	}
}

// Logout
// @Tags     Admin.Base
// @Summary  注销成功（jwt加入黑名单）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router   /admin-api/logout [delete]
func (b *Base) Logout(c context.Context, ctx *app.RequestContext) {
	token := ctx.Request.Header.Get("Authorization")
	jwtEntity := entities.JwtBlacklist{Jwt: token}
	if err := jwtService.Blacklist(jwtEntity); err != nil {
		global.LOG.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", ctx)
	} else {
		response.OkWithMessage("注销成功", ctx)
	}
}
