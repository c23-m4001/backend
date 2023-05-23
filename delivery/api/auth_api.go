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
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{token=dto_response.AuthTokenResponse}}
func (a *AuthApi) EmailLogin() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.AuthEmailLoginRequest
			ctx.mustBind(&request)

			token := a.authUseCase.LoginEmail(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"token": dto_response.NewAuthTokenResponse(token),
					},
				},
			)
		},
	)
}

//	@Router		/auth/google-login [post]
//	@Summary	Google Login
//	@tags		Auth
//	@Accept		json
//	@Param		dto_request.AuthGoogleLoginRequest	body	dto_request.AuthGoogleLoginRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{google_login_data=dto_response.GoogleLoginResponse}}
func (a *AuthApi) GoogleLogin() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.AuthGoogleLoginRequest
			ctx.mustBind(&request)

			googleLoginData := a.authUseCase.LoginGoogle(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"google_login_data": dto_response.NewGoogleLoginResponse(googleLoginData),
					},
				},
			)
		},
	)
}

//	@Router		/auth/email-register [post]
//	@Summary	Email Register
//	@tags		Auth
//	@Accept		json
//	@Param		dto_request.AuthEmailRegisterRequest	body	dto_request.AuthEmailRegisterRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{token=dto_response.AuthTokenResponse}}
func (a *AuthApi) EmailRegister() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.AuthEmailRegisterRequest
			ctx.mustBind(&request)

			token := a.authUseCase.RegisterEmail(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"token": dto_response.NewAuthTokenResponse(token),
					},
				},
			)
		},
	)
}

//	@Router		/auth/login-histories [post]
//	@Summary	Login Histories (up to 10 login histories)
//	@tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{login_histories=[]dto_response.LoginHistoryResponse}}
func (a *AuthApi) LoginHistory() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {

			userAccessTokens := a.authUseCase.LoginHistories(ctx.context())

			nodes := []dto_response.LoginHistoryResponse{}
			for _, userAccessToken := range userAccessTokens {
				nodes = append(nodes, dto_response.NewLoginHistoryResponse(userAccessToken))
			}

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"login_histories": nodes,
					},
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
	routerGroup.POST("/google-login", api.GoogleLogin())
	routerGroup.POST("/email-register", api.EmailRegister())
	routerGroup.POST("/login-histories", api.LoginHistory())
}
