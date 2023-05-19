package dto_request

import "capstone/data_type"

type WalletCreateRequest struct {
	Name     string                   `json:"name" validate:"required,not_empty" example:"Cash"`
	LogoType data_type.WalletLogoType `json:"logo_type" validate:"data_type_enum"`
} // @name WalletCreateRequest

type WalletFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name WalletFetchSorts

type WalletFetchRequest struct {
	PaginationRequest
	Sorts  WalletFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string          `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name WalletFetchRequest

type WalletGetRequest struct {
	WalletId string `json:"-" swaggerignore:"true"`
} // @name WalletGetRequest

type WalletUpdateRequest struct {
	Name     string                   `json:"name" validate:"required,not_empty" example:"Cash"`
	LogoType data_type.WalletLogoType `json:"logo_type" validate:"data_type_enum"`

	WalletId string `json:"-" swaggerignore:"true"`
} // @name WalletUpdateRequest
