package middleware

import (
	"capstone/internal/gin/validator"
	"capstone/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func TranslatorHandler(router gin.IRouter) {
	router.Use(func(ctx *gin.Context) {
		var (
			matcher = language.NewMatcher([]language.Tag{
				language.English, // The first language is used as fallback.
				language.Indonesian,
			})
			accept = ctx.GetHeader("Accept-Language")
		)

		tag, _ := language.MatchStrings(matcher, accept)
		locale := tag.String()

		if validatorTranslator, ok := validator.Translators[locale]; ok {
			ctx.Request = ctx.Request.WithContext(model.SetValidatorTranslatorCtx(ctx.Request.Context(), validatorTranslator))
		}

		ctx.Next()
	})
}
