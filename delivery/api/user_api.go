package api

import (
	"capstone/delivery/dto_response"
	"capstone/manager"
	"capstone/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	api

	authUseCase use_case.AuthUseCase
	userUseCase use_case.UserUseCase
}

//	@Router		/users/me [get]
//	@Summary	Get Current Logged In User Data
//	@tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserMeResponse}}
func (a *UserApi) Me() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			user := a.userUseCase.GetMe(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"user": dto_response.NewUserMeResponse(user),
					},
				},
			)
		},
	)
}

func RegisterUserApi(router gin.IRouter, useCaseManager manager.UseCaseManager) {
	api := UserApi{
		api:         newApi(),
		authUseCase: useCaseManager.AuthUseCase(),
		userUseCase: useCaseManager.UserUseCase(),
	}

	routerGroup := router.Group("/users")

	routerGroup.GET("/me", api.Me())
}
