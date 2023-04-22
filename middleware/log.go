package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/doublepi123/logs"
	"go.uber.org/zap"
)

func InitLog() {
	logs.InitLog("log/info.log", "log/err.log", zap.InfoLevel)
}

func Logger() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx = logs.NewContextWithLogId(ctx)

		logs.CtxInfo(ctx, "request access ", zap.String("uri", c.Request.URI().String()), zap.String("body", string(c.Request.BodyBytes())))

		c.Next(ctx)

		logs.CtxInfo(ctx, "request end", zap.String("resp", string(c.Response.BodyBytes())))
	}
}
