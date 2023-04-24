package middleware

import (
	"capstone/model"

	"github.com/gin-gonic/gin"
)

func IpHandler(router gin.IRouter) {
	router.Use(func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		// fmt.Println(ip, ctx.GetHeader("X-Forwarded-For"))
		ctx.Request = ctx.Request.WithContext(model.SetIpCtx(ctx.Request.Context(), ip))
		ctx.Next()
	})
}
