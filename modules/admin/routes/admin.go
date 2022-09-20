package routes

import (
	"HertzBoot/app/middleware"
	"HertzBoot/modules/admin/http/controllers/admin/v1"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitAdminRouter(Router *route.RouterGroup) {
	r := Router.Group("user").Use(middleware.Operations())
	noRecord := Router.Group("user")
	var baseApi = new(v1.Admin)
	{
		r.POST("register", baseApi.Register)                     // 用户注册账号
		r.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		r.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		r.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		r.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		r.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组

	}
	{
		noRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		noRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}
