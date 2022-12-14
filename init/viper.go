package init

import (
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Viper 解析配置文件优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(pkg.ConfigEnv); configEnv == "" {
				config = pkg.ConfigFile
				fmt.Printf("您正在使用CONFIG的默认值,config的路径为%v\n", pkg.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error from config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file content changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.CONFIG.JWT.ExpiresTime)),
	)
	return v
}
