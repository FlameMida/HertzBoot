package middleware

import (
	"HertzBoot/internal/core/entities"
	"HertzBoot/internal/core/http/requests"
	"HertzBoot/internal/core/service"
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"io"
	"strconv"
	"time"
)

var OperationsService = new(service.OperationsService)

func Operations() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var body []byte
		var userId int

		if string(ctx.Request.Method()) != consts.MethodGet {
			var err error
			body, err = io.ReadAll(ctx.Request.BodyStream())
			if err != nil {
				hlog.Error("read body from request error:", err.Error())
			} else {
				ctx.Request.SetBody(body)
			}
		}
		if claims, ok := ctx.Get("claims"); ok {
			waitUse := claims.(*requests.CustomClaims)
			userId = int(waitUse.ID)
		} else {
			id, err := strconv.Atoi(ctx.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			} else {
				userId = id
			}
		}
		record := entities.Operations{
			Ip:      ctx.ClientIP(),
			Method:  string(ctx.Request.Method()),
			Path:    string(ctx.Request.Path()),
			Agent:   string(ctx.UserAgent()),
			Body:    string(body),
			AdminID: userId,
		}
		writer := responseBodyWriter{
			ctx.GetWriter(),
			&bytes.Buffer{},
		}
		now := time.Now()
		ctx.Next(c)

		latency := time.Since(now)
		record.ErrorMessage = ctx.Errors.ByType(errors.ErrorTypePrivate).String()
		record.Status = ctx.Response.StatusCode()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := OperationsService.CreateOperations(record); err != nil {
			hlog.Error("create operation record error:", err.Error())
		}
	}
}

type responseBodyWriter struct {
	network.Writer
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.Write(b)
}
