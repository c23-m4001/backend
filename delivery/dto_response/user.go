package dto_response

import (
	"capstone/model"
)

type UserMeResponse struct {
	Id    string `json:"id" example:"023b6735-8255-43c0-bc3d-f6d1e423612d"`
	Email string `json:"email" example:"email@gmailcom"`
	Name  string `json:"name" example:"John Doe"`

	Wallets []WalletResponse `json:"wallets" extensions:"x-nullable"`
}

func NewUserMeResponse(user model.User) UserMeResponse {
	r := UserMeResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}

	if user.Wallets != nil {
		user.Wallets = []model.Wallet{}
		for _, wallet := range user.Wallets {
			r.Wallets = append(r.Wallets, NewWalletResponse(wallet))
		}
	}

	return r
}

func NewUserMeResponseP(user model.User) *UserMeResponse {
	r := NewUserMeResponse(user)
	return &r
}
