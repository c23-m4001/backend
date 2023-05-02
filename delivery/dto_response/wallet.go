package dto_response

import (
	"capstone/data_type"
	"capstone/model"
)

type WalletResponse struct {
	Id          string                   `json:"id" example:"023b6735-8255-43c0-bc3d-f6d1e423612d"`
	UserId      string                   `json:"user_id" example:"ccb77821-6289-468d-b4b2-a9b2efc60cb8"`
	Name        string                   `json:"name" example:"Cash"`
	TotalAmount float64                  `json:"total_amount" example:"10000"`
	LogoType    data_type.WalletLogoType `json:"logo_type"`
	Timestamp
} // @name WalletResponse

func NewWalletResponse(wallet model.Wallet) WalletResponse {
	r := WalletResponse{
		Id:          wallet.Id,
		UserId:      wallet.UserId,
		Name:        wallet.Name,
		TotalAmount: wallet.TotalAmount,
		LogoType:    wallet.LogoType,
		Timestamp:   Timestamp(wallet.Timestamp),
	}
	return r
}

func NewWalletResponseP(wallet model.Wallet) *WalletResponse {
	r := NewWalletResponse(wallet)
	return &r
}
