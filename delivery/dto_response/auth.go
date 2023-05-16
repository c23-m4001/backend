package dto_response

import (
	"capstone/data_type"
	"capstone/model"
)

type AuthTokenResponse struct {
	AccessToken          string             `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	AccessTokenExpiredAt data_type.DateTime `json:"access_token_expired_at" format:"YYYY-MM-DDTHH:mm:ssZ" example:"2006-01-02T15:04:05+07:00"`
	TokenType            string             `json:"token_type" example:"Bearer"`
} // @name AuthTokenResponse

func NewAuthTokenResponse(token model.Token) AuthTokenResponse {
	r := AuthTokenResponse{
		AccessToken:          token.AccessToken,
		AccessTokenExpiredAt: token.AccessTokenExpiredAt,
		TokenType:            token.TokenType,
	}
	return r
}

func NewAuthTokenResponseP(token model.Token) *AuthTokenResponse {
	r := NewAuthTokenResponse(token)
	return &r
}

type LoginHistoryResponse struct {
	Id           string             `json:"id" example:"284c93cb-1fed-4891-b5b8-032feb7c86da"`
	IpAddress    string             `json:"ip_address" example:"127.0.0.1"`
	LocationName string             `json:"location_name" example:"Medan North Sumatera"`
	Time         data_type.DateTime `json:"time" example:"2023-0514T14:33:26+07:00"`
} // @name LoginHistoryResponse

func NewLoginHistoryResponse(userAccessToken model.UserAccessToken) LoginHistoryResponse {
	r := LoginHistoryResponse{
		Id:           userAccessToken.Id,
		IpAddress:    *userAccessToken.IpAddress,
		LocationName: *userAccessToken.LocationName,
		Time:         userAccessToken.CreatedAt.DateTime(),
	}
	return r
}

func NewLoginHistoryResponseP(userAccessToken model.UserAccessToken) *LoginHistoryResponse {
	r := NewLoginHistoryResponse(userAccessToken)
	return &r
}

type GoogleLoginUserDataResponse struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"johndoe@gmail.com"`
}

type GoogleLoginResponse struct {
	UserData *GoogleLoginUserDataResponse `json:"user_data" extensions:"x-nullable"`
	Token    *AuthTokenResponse           `json:"token" extensions:"x-nullable"`
} // @name GoogleLoginResponse

func NewGoogleLoginResponse(googleLoginData model.GoogleLoginData) GoogleLoginResponse {
	r := GoogleLoginResponse{}

	if googleLoginData.UserData != nil {
		r.UserData = &GoogleLoginUserDataResponse{
			Name:  googleLoginData.UserData.Name,
			Email: googleLoginData.UserData.Email,
		}
	}

	if googleLoginData.Token != nil {
		r.Token = NewAuthTokenResponseP(*googleLoginData.Token)
	}

	return r
}

func NewGoogleLoginResponseP(googleLoginData model.GoogleLoginData) *GoogleLoginResponse {
	r := NewGoogleLoginResponse(googleLoginData)
	return &r
}
