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
//	@tags		Wallet
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
//	@Summary	Fetch current user wallet list
//	@tags		Wallet
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

//	@Router		/wallets/{id} [get]
//	@Summary	Get
//	@tags		Wallet
//	@Accept		json
//	@Produce	json
//	@Param		id								path		string							true	"Id"	format(uuid)
//	@Param		dto_request.WalletGetRequest	body		dto_request.WalletGetRequest	true	"Body Request"
//	@Success	200								{object}	dto_response.Response{data=dto_response.DataResponse{wallet=[]dto_response.WalletResponse}}
func (a *WalletApi) Get() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.WalletGetRequest

			request.WalletId = id

			wallet := a.walletUseCase.Get(ctx.context(), request)

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

//	@Router		/wallets/{id} [put]
//	@Summary	Update
//	@tags		Wallet
//	@Accept		json
//	@Produce	json
//	@Param		id								path		string							true	"Id"	format(uuid)
//	@Param		dto_request.WalletUpdateRequest	body		dto_request.WalletUpdateRequest	true	"Body Request"
//	@Success	200								{object}	dto_response.Response{data=dto_response.DataResponse{wallet=[]dto_response.WalletResponse}}
func (a *WalletApi) Update() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.WalletUpdateRequest
			ctx.mustBind(&request)

			request.WalletId = id

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

//	@Router		/wallets/{id} [delete]
//	@Summary	Delete wallet
//	@tags		Wallet
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Id"	format(uuid)
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *WalletApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.WalletDeleteRequest
			request.WalletId = id

			a.walletUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

//	@Router		/wallets/options/transaction-form [post]
//	@Summary	Option for Transaction Form
//	@tags		Wallet
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.WalletOptionForTransactionFormRequest	body		dto_request.WalletOptionForTransactionFormRequest	true	"Body Request"
//	@Success	200													{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.WalletResponse}}
func (a *WalletApi) OptionForTransactionForm() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.WalletOptionForTransactionFormRequest
			ctx.mustBind(&request)

			wallets, total := a.walletUseCase.OptionForTransactionForm(ctx.context(), request)

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

func RegisterWalletApi(router gin.IRouter, useCaseManager manager.UseCaseManager) {
	api := WalletApi{
		api:           newApi(),
		walletUseCase: useCaseManager.WalletUseCase(),
	}

	routerGroup := router.Group("/wallets")

	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/transaction-form", api.OptionForTransactionForm())
}
