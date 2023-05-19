package dto_response

import (
	"capstone/model"
)

type CategoryResponse struct {
	Id        string `json:"id" example:"023b6735-8255-43c0-bc3d-f6d1e423612d"`
	Name      string `json:"name" example:"Makanan"`
	IsGlobal  bool   `json:"is_global"`
	IsExpense bool   `json:"is_expense"`
	Timestamp
} // @name CategoryResponse

func NewCategoryResponse(category model.Category) CategoryResponse {
	r := CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		IsGlobal:  category.IsGlobal,
		IsExpense: category.IsExpense,
		Timestamp: Timestamp(category.Timestamp),
	}
	return r
}

func NewCategoryResponseP(category model.Category) *CategoryResponse {
	r := NewCategoryResponse(category)
	return &r
}
