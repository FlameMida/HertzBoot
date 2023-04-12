package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"time"
)

// Cors 处理跨域请求,支持options访问
func Cors() app.HandlerFunc {
	return cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD",
			},
			AllowHeaders: []string{
				"Content-Type",
				"Authorization",
				"X-Token",
			},
			ExposeHeaders: []string{
				"Content-Length",
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Headers",
				"Content-Type",
			},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})	
}
