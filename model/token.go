package model

import "capstone/data_type"

type Token struct {
	AccessToken          string
	AccessTokenExpiredAt data_type.DateTime
	TokenType            string
}
