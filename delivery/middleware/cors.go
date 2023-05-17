package middleware

import (
	"capstone/delivery/dto_response"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler(router gin.IRouter, allowedOrigins []string) {
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodHead,
		},
		AllowHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Authorization",
			"Cache-Control",
			"Content-Type",
			"Content-Length",
			"Origin",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           2 * time.Hour,
	})

	router.Use(func(ctx *gin.Context) {
		corsMiddleware(ctx)

		if ctx.Writer.Status() == http.StatusForbidden {
			panic(dto_response.NewForbiddenResponse("CORS Error: This endpoint does not allow requests from this origin"))
		}

		ctx.Next()
	})
}
