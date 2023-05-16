package jwt

import "time"

type Payload struct {
	Id        string
	UserId    string
	CreatedAt time.Time
	ExpiredAt time.Time
}
