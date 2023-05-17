package dto_request

type PaginationRequest struct {
	Page  *int `json:"page" validate:"required_with=Limit,omitempty,gte=1" example:"1" extensions:"x-nullable"`
	Limit *int `json:"limit" validate:"required_with=Page,omitempty,gte=1,lte=100" example:"100" extensions:"x-nullable"`
}
