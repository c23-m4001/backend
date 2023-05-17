package api

import (
	"capstone/delivery/dto_response"
	bindingInternal "capstone/internal/gin/binding"
	"capstone/internal/gin/validator"
	"capstone/model"
	"capstone/util"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
)

type apiContext struct {
	ginCtx *gin.Context
}

func newApiContext(ctx *gin.Context) apiContext {
	return apiContext{
		ginCtx: ctx,
	}
}

func (a *apiContext) context() context.Context {
	return a.ginCtx.Request.Context()
}

func (a *apiContext) getClientIp() string {
	return a.ginCtx.ClientIP()
}

func (a *apiContext) getParam(key string) string {
	return a.ginCtx.Param(key)
}

func (a *apiContext) getUuidParam(key string) string {
	uuidParam := a.getParam(key)

	if !util.IsUuid(uuidParam) {
		panic(dto_response.NewBadRequestResponse(fmt.Sprintf("%s must be a valid UUID", strcase.ToCamel(key))))
	}

	return uuidParam
}

func (a *apiContext) shouldBind(obj interface{}) error {
	return util.ShouldGinBind(a.ginCtx, obj)
}

func (a *apiContext) mustBind(obj interface{}) {
	if err := a.shouldBind(obj); err != nil {
		panic(a.translateBindErr(err))
	}
}

func (a *apiContext) translateBindErr(err error) dto_response.ErrorResponse {
	var r dto_response.ErrorResponse

	switch v := err.(type) {
	case validator.ValidationErrors:
		errs := []dto_response.Error{}
		translations := v.Translate(model.MustGetValidatorTranslatorCtx(a.context()))
		for k, translation := range translations {
			errs = append(errs, dto_response.Error{
				Domain:  k,
				Message: translation,
			})
		}

		r = dto_response.NewBadRequestResponse("Invalid request payload")
		r.Errors = errs

	case *json.UnmarshalTypeError:
		r = dto_response.NewBadRequestResponse("Invalid request payload (type error)")

	case *json.InvalidUnmarshalError:
		r = dto_response.NewBadRequestResponse("Invalid request payload (unmarshal error)")

	default:
		switch v {
		case bindingInternal.ErrConvertMapStringSlice,
			bindingInternal.ErrConvertToMapString,
			bindingInternal.ErrMultiFileHeader,
			bindingInternal.ErrMultiFileHeaderLenInvalid,
			bindingInternal.ErrIgnoredBinding:
			r = dto_response.NewBadRequestResponse("Invalid request payload")

		default:
			panic(err)
		}
	}

	return r
}

func (a *apiContext) json(code int, obj interface{}) {
	a.ginCtx.JSON(code, obj)
}
