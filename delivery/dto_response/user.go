package dto_response

import (
	"capstone/model"
)

type UserMeResponse struct {
	Id    string `json:"id" example:"023b6735-8255-43c0-bc3d-f6d1e423612d"`
	Email string `json:"email" example:"email@gmailcom"`
	Name  string `json:"name" example:"John Doe"`

	HaveWallet *bool `json:"have_wallet" extensions:"x-nullable"`
} // @name UserMeResponse

func NewUserMeResponse(user model.User) UserMeResponse {
	r := UserMeResponse{
		Id:         user.Id,
		Email:      user.Email,
		Name:       user.Name,
		HaveWallet: user.HaveWallet,
	}
	return r
}

func NewUserMeResponseP(user model.User) *UserMeResponse {
	r := NewUserMeResponse(user)
	return &r
}
