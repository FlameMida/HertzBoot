package routes

import (
	"HertzBoot/internal/core/http/controllers/admin/v1"
	"github.com/cloudwego/hertz/pkg/route"
)

func AdminInitDBRouter(Router *route.RouterGroup) {
	r := Router.Group("init")
	var dbApi = new(v1.DB)
	{
		r.POST("initDB", dbApi.InitDB)
		r.POST("checkDB", dbApi.CheckDB)
	}
}
