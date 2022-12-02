package service

import (
	"HertzBoot/internal/core/entities"
	"HertzBoot/pkg/global"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

type JwtService struct {
}

// @author:      Flame
// @function:    Blacklist
// @description: 拉黑jwt
// @param:       jwtList model.JwtBlacklist
// @return:      err error

func (jwtService *JwtService) Blacklist(jwtList entities.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// @author:      Flame
// @function:    IsBlacklist
// @description: 判断JWT是否在黑名单内部
// @param:       jwt string
// @return:      bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// @author:      Flame
// @function:    GetRedisJWT
// @description: 从redis取jwt
// @param:       userName string
// @return:      err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

// @author:      Flame
// @function:    SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param:       jwt string, userName string
// @return:      err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.DB.Model(&entities.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		hlog.Error("加载数据库jwt黑名单失败!", err.Error())
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
