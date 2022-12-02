package global

import (
	"HertzBoot/pkg/timer"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"golang.org/x/sync/singleflight"

	"HertzBoot/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	REDIS              *redis.Client
	CONFIG             config.Server
	VP                 *viper.Viper
	Timer              = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{}
	BlackCache         local_cache.Cache
)
