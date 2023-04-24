package model

import (
	"context"
)

type ipCtxKeyType string

var ipCtxKey = ipCtxKeyType("capstone-ip")

func SetIpCtx(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, ipCtxKey, ip)
}

func GetIpCtx(ctx context.Context) string {
	v, _ := ctx.Value(ipCtxKey).(string)

	return v
}
