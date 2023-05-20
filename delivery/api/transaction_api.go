package api

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/manager"
	"capstone/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionApi struct {
	api

	transactionUseCase use_case.TransactionUseCase
}

//	@Router		/transactions [post]
//	@Summary	Create
//	@tags		Transaction
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.TransactionCreateRequest	body		dto_request.TransactionCreateRequest	true	"Body Request"
//	@Success	200										{object}	dto_response.Response{data=dto_response.DataResponse{transaction=dto_response.TransactionResponse}}
func (a *TransactionApi) Create() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.TransactionCreateRequest
			ctx.mustBind(&request)

			transaction := a.transactionUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"transaction": dto_response.NewTransactionResponse(transaction),
					},
				},
			)
		},
	)
}

//	@Router		/transactions/filter [post]
//	@Summary	Fetch current user transaction list
//	@tags		Transaction
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.TransactionFetchRequest	body		dto_request.TransactionFetchRequest	true	"Body Request"
//	@Success	200									{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.TransactionResponse}}
func (a *TransactionApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.TransactionFetchRequest
			ctx.mustBind(&request)

			transactions, total := a.transactionUseCase.Fetch(ctx.context(), request)

			nodes := []dto_response.TransactionResponse{}
			for _, transaction := range transactions {
				nodes = append(nodes, dto_response.NewTransactionResponse(transaction))
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

//	@Router		/transactions/{id} [get]
//	@Summary	Get
//	@tags		Transaction
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.TransactionGetRequest	body		dto_request.TransactionGetRequest	true	"Body Request"
//	@Param		id									path		string								true	"Id"	format(uuid)
//	@Success	200									{object}	dto_response.Response{data=dto_response.DataResponse{transaction=[]dto_response.TransactionResponse}}
func (a *TransactionApi) Get() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.TransactionGetRequest

			request.TransactionId = id

			transaction := a.transactionUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"transaction": dto_response.NewTransactionResponse(transaction),
					},
				},
			)
		},
	)
}

//	@Router		/transactions/{id} [put]
//	@Summary	Update
//	@tags		Transaction
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.TransactionUpdateRequest	body		dto_request.TransactionUpdateRequest	true	"Body Request"
//	@Param		id										path		string									true	"Id"	format(uuid)
//	@Success	200										{object}	dto_response.Response{data=dto_response.DataResponse{transaction=[]dto_response.TransactionResponse}}
func (a *TransactionApi) Update() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.TransactionUpdateRequest
			ctx.mustBind(&request)

			request.TransactionId = id

			transaction := a.transactionUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"transaction": dto_response.NewTransactionResponse(transaction),
					},
				},
			)
		},
	)
}

//	@Router		/transactions/{id} [delete]
//	@Summary	Delete
//	@tags		Transaction
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Id"	format(uuid)
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *TransactionApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.TransactionDeleteRequest
			ctx.mustBind(&request)

			request.TransactionId = id

			a.transactionUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterTransactionApi(router gin.IRouter, useCaseManager manager.UseCaseManager) {
	api := TransactionApi{
		api:                newApi(),
		transactionUseCase: useCaseManager.TransactionUseCase(),
	}

	routerGroup := router.Group("/transactions")

	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}
