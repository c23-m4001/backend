package dto_request

import (
	"capstone/data_type"
)

type TransactionCreateRequest struct {
	CategoryId string         `json:"category_id" validate:"required,not_empty,uuid" example:"78cfc498-7e4e-4f91-aaa7-6ada8525d68c"`
	WalletId   string         `json:"wallet_id" validate:"required,not_empty,uuid" example:"ae462e92-a133-4b93-bba6-7fd1ec037f33"`
	Name       string         `json:"name" validate:"required,not_empty" example:"Makan Siang"`
	Amount     float64        `json:"amount" validate:"required,gt=0" example:"10000"`
	Date       data_type.Date `json:"date" format:"YYYY-MM-DD" example:"2006-01-02"`
} // @name TransactionCreateRequest

type TransactionFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name date" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name TransactionFetchSorts

type TransactionFetchRequest struct {
	PaginationRequest
	Sorts TransactionFetchSorts `json:"sorts" validate:"unique=Field,dive"`

	StartDate  data_type.Date `json:"start_date" format:"YYYY-MM-DD" example:"2006-01-02"`
	EndDate    data_type.Date `json:"end_date" format:"YYYY-MM-DD" example:"2006-01-02"`
	CategoryId *string        `json:"category_id" validate:"omitempty,not_empty,uuid" example:"116e6126-77ce-45ab-8e4d-f7cc2b00cccf" extensions:"x-nullable"`
	WalletId   *string        `json:"wallet_id" validate:"omitempty,not_empty,uuid" example:"116e6126-77ce-45ab-8e4d-f7cc2b00cccf" extensions:"x-nullable"`

	Phrase *string `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name TransactionFetchRequest

type TransactionGetSummaryTotalRequest struct {
	CategoryId *string `json:"category_id" validate:"omitempty,not_empty,uuid" example:"116e6126-77ce-45ab-8e4d-f7cc2b00cccf" extensions:"x-nullable"`
	WalletId   *string `json:"wallet_id" validate:"omitempty,not_empty,uuid" example:"116e6126-77ce-45ab-8e4d-f7cc2b00cccf" extensions:"x-nullable"`
} // @name TransactionGetSummaryTotalRequest

type TransactionGetSummaryRequest struct {
	StartDate  data_type.Date `json:"start_date" format:"YYYY-MM-DD" example:"2006-01-02"`
	EndDate    data_type.Date `json:"end_date" format:"YYYY-MM-DD" example:"2006-01-02"`
	CategoryId *string        `json:"category_id" validate:"omitempty,not_empty,uuid" example:"116e6126-77ce-45ab-8e4d-f7cc2b00cccf" extensions:"x-nullable"`
	WalletId   *string        `json:"wallet_id" validate:"omitempty,not_empty,uuid" example:"116e6126-77ce-45ab-8e4d-f7cc2b00cccf" extensions:"x-nullable"`
} // @name TransactionGetSummaryRequest

type TransactionGetRequest struct {
	TransactionId string `json:"-" swaggerignore:"true"`
} // @name TransactionGetRequest

type TransactionUpdateRequest struct {
	Name   string         `json:"name" validate:"required,not_empty" example:"Makan Siang"`
	Amount float64        `json:"amount" validate:"required,gte=0" example:"10000"`
	Date   data_type.Date `json:"date" format:"YYYY-MM-DD" example:"2006-01-02"`

	TransactionId string `json:"-" swaggerignore:"true"`
} // @name TransactionUpdateRequest

type TransactionDeleteRequest struct {
	TransactionId string `json:"-" swaggerignore:"true"`
} // @name TransactionDeleteRequest
