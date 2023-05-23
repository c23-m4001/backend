package dto_response

import (
	"capstone/data_type"
	"capstone/model"
)

type TransactionResponse struct {
	Id         string         `json:"id" example:"4aa51ced-c7fe-4f8a-aca4-1b4ccaed00cc"`
	CategoryId string         `json:"category_id" example:"05442cc8-b88c-4005-a38e-36135ad2f41c"`
	WalletId   string         `json:"wallet_id" example:"a5ad60f0-6efd-47d3-a166-20c9c00f75ed"`
	Name       string         `json:"name" example:"Makan Siang"`
	Amount     float64        `json:"amount" example:"10000"`
	Date       data_type.Date `json:"date" example:"date"`

	Category *CategoryResponse `json:"category" extensions:"x-nullable"`
} // @name TransactionResponse

func NewTransactionResponse(transaction model.Transaction) TransactionResponse {
	r := TransactionResponse{
		Id:         transaction.Id,
		CategoryId: transaction.CategoryId,
		WalletId:   transaction.WalletId,
		Name:       transaction.Name,
		Amount:     transaction.Amount,
		Date:       transaction.Date,
	}

	if transaction.Category != nil {
		r.Category = NewCategoryResponseP(*transaction.Category)
	}

	return r
}

func NewTransactionResponseP(transaction model.Transaction) *TransactionResponse {
	r := NewTransactionResponse(transaction)
	return &r
}
