package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	jwtLib "github.com/golang-jwt/jwt/v4"
)

func mustGetPrivateKey(privateKeyFilePath string) *rsa.PrivateKey {
	privateKey, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Errorf("private key not found in %s", privateKeyFilePath))
		}

		panic(err)
	}

	block, _ := pem.Decode(privateKey)
	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return rsaPrivateKey
}

func mustGetPublicKey(publicKeyPath string) *rsa.PublicKey {
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Errorf("public key not found in %s", publicKeyPath))
		}

		panic(err)
	}

	block, _ := pem.Decode(publicKey)
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return rsaPublicKey
}

type jwt struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func (j *jwt) signingMethod() jwtLib.SigningMethod {
	return jwtLib.SigningMethodRS256
}

func (j *jwt) tokenType() string {
	return "Bearer"
}

func (j *jwt) constructToken(signedToken string) *Token {
	return &Token{
		Type:        j.tokenType(),
		AccessToken: signedToken,
	}
}

func (j *jwt) parseToken(finalToken string) (string, error) {
	token, err := parseToken(finalToken)
	if err != nil {
		return "", ErrInvalidToken
	}

	if token.Type != j.tokenType() {
		return "", ErrInvalidToken
	}

	return token.AccessToken, nil
}

func (j *jwt) Generate(payload Payload) (*Token, error) {
	token := jwtLib.NewWithClaims(j.signingMethod(), jwtLib.RegisteredClaims{
		Audience:  []string{payload.UserId},
		ExpiresAt: &jwtLib.NumericDate{Time: payload.ExpiredAt},
		IssuedAt:  &jwtLib.NumericDate{Time: payload.CreatedAt},
		NotBefore: &jwtLib.NumericDate{Time: payload.CreatedAt},
		ID:        payload.Id,
		Subject:   payload.UserId,
	})

	signedToken, err := token.SignedString(j.privateKey)
	if err != nil {
		return nil, err
	}

	finalizedToken := j.constructToken(signedToken)
	finalizedToken.ExpiredAt = payload.ExpiredAt

	return finalizedToken, nil
}

func (j *jwt) Parse(finalToken string) (*Payload, error) {
	signedToken, err := j.parseToken(finalToken)
	if err != nil {
		return nil, err
	}

	claims := jwtLib.RegisteredClaims{}

	_, err = jwtLib.ParseWithClaims(signedToken, &claims, func(t *jwtLib.Token) (interface{}, error) {
		if t.Method != j.signingMethod() {
			return nil, ErrInvalidToken
		}

		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	payload := Payload{
		Id:        claims.ID,
		UserId:    claims.Subject,
		CreatedAt: claims.IssuedAt.Time,
		ExpiredAt: claims.ExpiresAt.Time,
	}

	return &payload, nil
}

func NewJwt(privateKeyPath string, publicKeyPath string) Jwt {
	return &jwt{
		privateKey: mustGetPrivateKey(privateKeyPath),
		publicKey:  mustGetPublicKey(publicKeyPath),
	}
}
