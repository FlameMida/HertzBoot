package provider

import (
	"HertzBoot/app/global"
	"HertzBoot/modules/core/service"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func RunServer() {
	// 从db加载jwt数据
	if global.DB != nil {
		service.LoadAll()
	}
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	Router := Routers(address)

	fmt.Printf("[Hertz-STARTER]文档地址:http://127.0.0.1%s/swagger/index.html \n", address)
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))
	Router.Spin()

}
