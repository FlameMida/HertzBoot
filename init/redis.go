package init

import (
	"HertzBoot/pkg/global"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	redisConfig := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		hlog.Error("redis connect ping failed, err:", err.Error())
	} else {
		hlog.Info("redis connect ping response: ", pong)
		global.REDIS = client
	}
}
