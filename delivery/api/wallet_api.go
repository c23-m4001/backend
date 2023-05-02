package api

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/manager"
	"capstone/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletApi struct {
	api

	walletUseCase use_case.WalletUseCase
}

//	@Router		/wallets [post]
//	@Summary	Create
//	@tags		User
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.WalletCreateRequest	body		dto_request.WalletCreateRequest	true	"Body Request"
//	@Success	200								{object}	dto_response.Response{data=dto_response.DataResponse{wallet=dto_response.WalletResponse}}
func (a *WalletApi) Create() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.WalletCreateRequest
			ctx.mustBind(&request)

			wallet := a.walletUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"wallet": dto_response.NewWalletResponse(wallet),
					},
				},
			)
		},
	)
}

//	@Router		/wallets/filter [post]
//	@Summary	Fetch
//	@tags		User
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.WalletFetchRequest	body		dto_request.WalletFetchRequest	true	"Body Request"
//	@Success	200								{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.WalletResponse}}
func (a *WalletApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.WalletFetchRequest
			ctx.mustBind(&request)

			wallets, total := a.walletUseCase.Fetch(ctx.context(), request)

			nodes := []dto_response.WalletResponse{}
			for _, wallet := range wallets {
				nodes = append(nodes, dto_response.NewWalletResponse(wallet))
			}

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Limit: request.Limit,
						Page:  request.Page,
						Nodes: nodes,
						Total: total,
					},
				},
			)
		},
	)
}

//	@Router		/wallets/{id} [put]
//	@Summary	Update
//	@tags		User
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.WalletUpdateRequest	body		dto_request.WalletUpdateRequest	true	"Body Request"
//	@Param		id								path		string							true	"Id"	format(uuid)
//	@Success	200								{object}	dto_response.Response{data=dto_response.DataResponse{wallet=[]dto_response.WalletResponse}}
func (a *WalletApi) Update() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.WalletUpdateRequest
			ctx.mustBind(&request)

			wallet := a.walletUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"wallet": dto_response.NewWalletResponse(wallet),
					},
				},
			)
		},
	)
}

func RegisterWalletApi(router gin.IRouter, useCaseManager manager.UseCaseManager) {
	api := WalletApi{
		api:           newApi(),
		walletUseCase: useCaseManager.WalletUseCase(),
	}

	routerGroup := router.Group("/wallets")

	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.PUT("/:id", api.Update())
}
