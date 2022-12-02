package init

import (
	"HertzBoot/internal/core/service"
	"HertzBoot/pkg/global"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

func RunServer() {
	// 从db加载jwt数据
	if global.DB != nil {
		service.LoadAll()
	}
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	Router := Routers(address)

	fmt.Printf("[HertzBoot]文档地址:http://127.0.0.1%s/swagger/index.html \n", address)
	time.Sleep(10 * time.Microsecond)
	hlog.Info("server run success on: ", address)
	Router.Spin()

}
