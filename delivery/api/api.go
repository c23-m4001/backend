package api

import (
	"capstone/config"
	"capstone/delivery/middleware"
	"capstone/manager"
	"capstone/model"

	"github.com/gin-gonic/gin"
)

type api struct {
}

func (a *api) Authorize(fn func(ctx apiContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiCtx := newApiContext(ctx)

		// check user authenticated
		model.MustGetUserCtx(apiCtx.context())

		fn(apiCtx)
	}
}

func (a *api) Guest(fn func(ctx apiContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(newApiContext(ctx))
	}
}

func newApi() api {
	return api{}
}

func registerMiddlewares(router gin.IRouter, container *manager.Container) {
	useCaseManager := container.UseCaseManager()
	loggerStack := container.InfrastructureManager().GetLoggerStack()

	middleware.TranslatorHandler(router)
	middleware.PanicHandler(router, loggerStack)
	middleware.CorsHandler(router, config.GetConfig().CorsAllowedOrigins)
	middleware.IpHandler(router)
	middleware.JWTHandler(router, useCaseManager.AuthUseCase())
}

func registerRoutes(router gin.IRouter, useCasemanager manager.UseCaseManager) {
	RegisterAuthApi(router, useCasemanager)
	RegisterCategoryApi(router, useCasemanager)
	RegisterUserApi(router, useCasemanager)
	RegisterWalletApi(router, useCasemanager)
}

func NewRouter(container *manager.Container) *gin.Engine {
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	registerMiddlewares(router, container)

	registerRoutes(router, container.UseCaseManager())

	return router
}
