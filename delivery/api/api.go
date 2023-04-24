package api

import (
	"capstone/config"
	"capstone/delivery/middleware"
	"capstone/manager"
	"capstone/model"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	middleware.IpHandler(router)
	middleware.JWTHandler(router, useCaseManager.AuthUseCase())
}

func registerRoutes(router gin.IRouter, useCasemanager manager.UseCaseManager) {
	RegisterAuthApi(router, useCasemanager)
	RegisterUserApi(router, useCasemanager)
}

func NewRouter(container *manager.Container) *gin.Engine {
	allowedHeaders := []string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Content-Length",
		"Origin",
		"X-CSRF-Token",
		"X-Requested-With",
	}

	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.GetConfig().CorsAllowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodHead,
		},
		AllowHeaders: allowedHeaders,
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           2 * time.Hour,
	}))

	registerMiddlewares(router, container)

	registerRoutes(router, container.UseCaseManager())

	return router
}
