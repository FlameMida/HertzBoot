package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 10001
	SUCCESS = 0
)

func Result(code int, data any, msg string, ctx *app.RequestContext) {
	// 开始时间
	ctx.JSON(consts.StatusOK, Response{
		code,
		data,
		msg,
	})
	return
}

func Ok(ctx *app.RequestContext) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", ctx)
}

func OkWithMessage(message string, ctx *app.RequestContext) {
	Result(SUCCESS, map[string]interface{}{}, message, ctx)
}

func OkWithData(data any, ctx *app.RequestContext) {
	Result(SUCCESS, data, "操作成功", ctx)
}

func OkWithDetailed(data any, message string, ctx *app.RequestContext) {
	Result(SUCCESS, data, message, ctx)
}

func Fail(ctx *app.RequestContext) {
	Result(ERROR, map[string]interface{}{}, "操作失败", ctx)
}

func FailWithMessage(message string, ctx *app.RequestContext) {
	Result(ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data any, message string, ctx *app.RequestContext) {
	Result(ERROR, data, message, ctx)
}
