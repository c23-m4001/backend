package util

import (
	bindingInternal "capstone/internal/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ShouldGinBind(ctx *gin.Context, obj interface{}) error {
	var (
		err error
	)

	b := bindingInternal.Default(ctx.Request.Method, ctx.ContentType())

	bb, ok := b.(binding.BindingBody)
	if ok {
		err = ctx.ShouldBindBodyWith(obj, bb)
	} else {
		err = ctx.ShouldBindWith(obj, b)
	}

	return err
}
