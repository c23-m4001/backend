package dto_request

import (
	"capstone/data_type"

	"github.com/go-playground/validator/v10"
)

type PaginationRequest struct {
	Page  *int `json:"page" validate:"required_with=Limit,omitempty,gte=1" example:"1" extensions:"x-nullable"`
	Limit *int `json:"limit" validate:"required_with=Page,omitempty,gte=1,lte=100" example:"100" extensions:"x-nullable"`
}

func DateRangeRequestValidationFn(sl validator.StructLevel) {
	var startDate, endDate data_type.Date
	switch v := sl.Current().Interface().(type) {
	case TransactionFetchRequest:
		startDate = v.StartDate
		endDate = v.EndDate
	}

	if startDate.IsValid() && endDate.IsValid() && endDate.IsLessThan(startDate) {
		sl.ReportError(endDate, "EndDate", "EndDate", "gtefield", "StartDate")
	}
}

func NullDateRangeRequestValidationFn(sl validator.StructLevel) {
	var startNullDate, endNullDate data_type.NullDate
	switch v := sl.Current().Interface().(type) {
	default:
		_ = v
	}

	startDate := startNullDate.DateP()
	endDate := endNullDate.DateP()

	if startDate != nil && endDate != nil && startDate.IsValid() && endDate.IsValid() && endDate.IsLessThan(*startDate) {
		sl.ReportError(endDate, "EndDate", "EndDate", "gtefield", "StartDate")
	}
}
