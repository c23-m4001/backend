package middleware

import (
	"capstone/constant"
	"capstone/model"
	"capstone/use_case"

	"github.com/gin-gonic/gin"
)

func JWTHandler(router gin.IRouter, authUseCase use_case.AuthUseCase) {
	router.Use(func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.Next()
			return
		}

		user, err := authUseCase.Parse(ctx, token)
		if err != nil {
			if err != constant.ErrNotAuthenticated {
				panic(err)
			}

			ctx.Next()
			return
		}

		ctx.Request = ctx.Request.WithContext(model.SetUserCtx(ctx.Request.Context(), user))
		ctx.Next()
	})
}
