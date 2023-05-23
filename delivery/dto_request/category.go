package dto_request

import "capstone/data_type"

type CategoryCreateRequest struct {
	Name      string                     `json:"name" validate:"required,not_empty" example:"Makanan"`
	IsExpense bool                       `json:"is_expense"`
	LogoType  data_type.CategoryLogoType `json:"logo_type" validate:"required,data_type_enum"`
} // @name CategoryCreateRequest

type CategoryFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CategoryFetchSorts

type CategoryFetchRequest struct {
	PaginationRequest
	Sorts     CategoryFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	IsExpense *bool              `json:"is_expense"`
	IsGlobal  *bool              `json:"is_global"`
	Phrase    *string            `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CategoryFetchRequest

type CategoryGetRequest struct {
	CategoryId string `json:"-" swaggerignore:"true"`
} // @name CategoryGetRequest

type CategoryUpdateRequest struct {
	Name      string                     `json:"name" validate:"required,not_empty"`
	IsExpense bool                       `json:"is_expense"`
	LogoType  data_type.CategoryLogoType `json:"logo_type" validate:"required,data_type_enum"`

	CategoryId string `json:"-" swaggerignore:"true"`
} // @name CategoryUpdateRequest

type CategoryDeleteRequest struct {
	CategoryId string `json:"-" swaggerignore:"true"`
} // @name CategoryDeleteRequest

type CategoryOptionForTransactionFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CategoryOptionForTransactionFormSorts

type CategoryOptionForTransactionFormRequest struct {
	PaginationRequest
	Sorts     CategoryOptionForTransactionFormSorts `json:"sorts" validate:"unique=Field,dive"`
	IsExpense *bool                                 `json:"is_expense"`
	Phrase    *string                               `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CategoryOptionForTransactionFormRequest
