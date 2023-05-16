package jwt

type Jwt interface {
	Generate(payload Payload) (*Token, error)
	Parse(finalToken string) (*Payload, error)
}
