package model

import (
	"capstone/internal/gin/validator"
	"context"

	ut "github.com/go-playground/universal-translator"
)

type validatorTranslatorCtxKeyType string

var validatorTranslatorCtxKey = validatorTranslatorCtxKeyType("capstone-validator-translator")

func SetValidatorTranslatorCtx(ctx context.Context, validatorTranslator ut.Translator) context.Context {
	return context.WithValue(ctx, validatorTranslatorCtxKey, validatorTranslator)
}

func MustGetValidatorTranslatorCtx(ctx context.Context) ut.Translator {
	v, ok := ctx.Value(validatorTranslatorCtxKey).(ut.Translator)
	if !ok {
		return validator.DefaultTranslator
	}

	return v
}
