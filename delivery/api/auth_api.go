package api

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/manager"
	"capstone/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	api
	authUseCase use_case.AuthUseCase
}

//	@Router		/auth/email-login [post]
//	@Summary	Email Login
//	@tags		Auth
//	@Accept		json
//	@Param		dto_request.AuthEmailLoginRequest	body	dto_request.AuthEmailLoginRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.AuthTokenResponse}
func (a *AuthApi) EmailLogin() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.AuthEmailLoginRequest
			ctx.mustBind(&request)

			data := a.authUseCase.LoginEmail(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.NewAuthTokenResponse(data),
				},
			)
		},
	)
}

func RegisterAuthApi(router gin.IRouter, useCaseManager manager.UseCaseManager) {
	api := AuthApi{
		api:         newApi(),
		authUseCase: useCaseManager.AuthUseCase(),
	}

	routerGroup := router.Group("/auth")

	routerGroup.POST("/email-login", api.EmailLogin())
}
