package api

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/manager"
	"capstone/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct {
	api

	categoryUseCase use_case.CategoryUseCase
}

//	@Router		/categories [post]
//	@Summary	Create
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.CategoryCreateRequest	body		dto_request.CategoryCreateRequest	true	"Body Request"
//	@Success	200									{object}	dto_response.Response{data=dto_response.DataResponse{category=dto_response.CategoryResponse}}
func (a *CategoryApi) Create() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.CategoryCreateRequest
			ctx.mustBind(&request)

			category := a.categoryUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"category": dto_response.NewCategoryResponse(category),
					},
				},
			)
		},
	)
}

//	@Router		/categories/filter [post]
//	@Summary	Fetch current user category list
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.CategoryFetchRequest	body		dto_request.CategoryFetchRequest	true	"Body Request"
//	@Success	200									{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CategoryResponse}}
func (a *CategoryApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.CategoryFetchRequest
			ctx.mustBind(&request)

			categories, total := a.categoryUseCase.Fetch(ctx.context(), request)

			nodes := []dto_response.CategoryResponse{}
			for _, category := range categories {
				nodes = append(nodes, dto_response.NewCategoryResponse(category))
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

//	@Router		/categories/defaults [post]
//	@Summary	Fetch Default Categories
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{categories=[]dto_response.CategoryResponse}}
func (a *CategoryApi) FetchDefaults() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			categories := a.categoryUseCase.FetchDefaults(ctx.context())

			nodes := []dto_response.CategoryResponse{}
			for _, category := range categories {
				nodes = append(nodes, dto_response.NewCategoryResponse(category))
			}

			ctx.json(
				http.StatusOK,
				dto_response.DataResponse{
					"categories": nodes,
				},
			)
		},
	)
}

//	@Router		/categories/{id} [get]
//	@Summary	Get
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.CategoryGetRequest	body		dto_request.CategoryGetRequest	true	"Body Request"
//	@Param		id								path		string							true	"Id"	format(uuid)
//	@Success	200								{object}	dto_response.Response{data=dto_response.DataResponse{category=[]dto_response.CategoryResponse}}
func (a *CategoryApi) Get() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.CategoryGetRequest

			request.CategoryId = id

			category := a.categoryUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"category": dto_response.NewCategoryResponse(category),
					},
				},
			)
		},
	)
}

//	@Router		/categories/{id} [put]
//	@Summary	Update
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.CategoryUpdateRequest	body		dto_request.CategoryUpdateRequest	true	"Body Request"
//	@Param		id									path		string								true	"Id"	format(uuid)
//	@Success	200									{object}	dto_response.Response{data=dto_response.DataResponse{category=[]dto_response.CategoryResponse}}
func (a *CategoryApi) Update() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.CategoryUpdateRequest
			ctx.mustBind(&request)

			request.CategoryId = id

			category := a.categoryUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"category": dto_response.NewCategoryResponse(category),
					},
				},
			)
		},
	)
}

//	@Router		/categories/{id} [delete]
//	@Summary	Delete
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.CategoryDeleteRequest	body		dto_request.CategoryDeleteRequest	true	"Body Request"
//	@Param		id									path		string								true	"Id"	format(uuid)
//	@Success	200									{object}	dto_response.SuccessResponse
func (a *CategoryApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.CategoryDeleteRequest
			ctx.mustBind(&request)

			request.CategoryId = id

			a.categoryUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

//	@Router		/categories/options/transaction-form [post]
//	@Summary	Option for Transaction Form
//	@tags		Category
//	@Accept		json
//	@Produce	json
//	@Param		dto_request.CategoryOptionForTransactionFormRequest	body		dto_request.CategoryOptionForTransactionFormRequest	true	"Body Request"
//	@Success	200													{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CategoryResponse}}
func (a *CategoryApi) OptionForTransactionForm() gin.HandlerFunc {
	return a.Authorize(
		func(ctx apiContext) {
			var request dto_request.CategoryOptionForTransactionFormRequest
			ctx.mustBind(&request)

			categories, total := a.categoryUseCase.OptionForTransactionForm(ctx.context(), request)

			nodes := []dto_response.CategoryResponse{}
			for _, category := range categories {
				nodes = append(nodes, dto_response.NewCategoryResponse(category))
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

func RegisterCategoryApi(router gin.IRouter, useCaseManager manager.UseCaseManager) {
	api := CategoryApi{
		api:             newApi(),
		categoryUseCase: useCaseManager.CategoryUseCase(),
	}

	routerGroup := router.Group("/categories")

	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.POST("/defaults", api.FetchDefaults())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/transaction-form", api.OptionForTransactionForm())
}
