package dto_request

type CategoryCreateRequest struct {
	Name      string `json:"name" validate:"required,not_empty" example:"Makanan"`
	IsExpense bool   `json:"is_expense"`
} // @name CategoryCreateRequest

type CategoryFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CategoryFetchSorts

type CategoryFetchRequest struct {
	PaginationRequest
	Sorts  CategoryFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string            `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CategoryFetchRequest

type CategoryGetRequest struct {
	CategoryId string `json:"-" swaggerignore:"true"`
} // @name CategoryGetRequest

type CategoryUpdateRequest struct {
	Name      string `json:"name" validate:"required,not_empty"`
	IsExpense bool   `json:"is_expense"`

	CategoryId string `json:"-" swaggerignore:"true"`
} // @name CategoryUpdateRequest

type CategoryDeleteRequest struct {
	CategoryId string `json:"-" swaggerignore:"true"`
} // @name CategoryDeleteRequest
