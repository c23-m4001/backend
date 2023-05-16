package jwt

import (
	"strings"
	"time"
)

type Token struct {
	Type        string
	AccessToken string
	ExpiredAt   time.Time
}

func parseToken(token string) (*Token, error) {
	split := strings.Split(token, " ")
	if len(split) != 2 {
		return nil, ErrInvalidToken
	}

	return &Token{
		Type:        split[0],
		AccessToken: split[1],
	}, nil
}
