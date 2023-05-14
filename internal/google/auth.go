package google

import (
	"capstone/config"
	"errors"
	"time"

	jwtLib "github.com/golang-jwt/jwt/v4"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const URL_ACCESS_TOKEN = "https://oauth2.googleapis.com/token"
const URL_REVOKE_ACCESS_TOKEN = "https://oauth2.googleapis.com/revoke"

type GoogleAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	IDToken     string `json:"id_token"`
}

func (r *GoogleAccessTokenResponse) RevokeAccessToken() error {
	data := url.Values{}
	data.Set("token", r.AccessToken)

	response, err := http.Post(URL_REVOKE_ACCESS_TOKEN, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var revokeErrorResponse struct {
		Error string `json:"error"`
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &revokeErrorResponse); err != nil {
		return err
	}

	if revokeErrorResponse.Error != "" {
		return errors.New(revokeErrorResponse.Error)
	}

	return nil
}

type GoogleJwtPayload struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
}

func (r GoogleJwtPayload) Valid() error {
	if r.Iss != "https://accounts.google.com" && r.Iss != "accounts.google.com" {
		return jwtLib.NewValidationError("Invalid token issuer", jwtLib.ValidationErrorIssuer)
	}

	if r.Aud != config.GetGoogleOauthConfig().ClientId {
		return jwtLib.NewValidationError("Invalid audience", jwtLib.ValidationErrorAudience)
	}

	if time.Now().Unix() >= int64(r.Exp) {
		return jwtLib.NewValidationError("Token Expired", jwtLib.ValidationErrorExpired)
	}

	return nil
}

type GoogleUserDataResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type GoogleErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func GoogleGetUserDataFromCode(code string) (*GoogleJwtPayload, *GoogleErrorResponse, error) {
	googleOauthConfig := config.GetGoogleOauthConfig()

	data := url.Values{}
	data.Set("client_id", googleOauthConfig.ClientId)
	data.Set("client_secret", googleOauthConfig.ClientSecret)
	data.Set("redirect_uri", googleOauthConfig.RedirectUri)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")

	response, err := http.Post(URL_ACCESS_TOKEN, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	var dataResponse GoogleAccessTokenResponse
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return nil, nil, err
	}

	if dataResponse.AccessToken == "" {
		var errorResponse GoogleErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return nil, nil, err
		}
		return nil, &errorResponse, nil
	}

	err = dataResponse.RevokeAccessToken()
	if err != nil {
		panic(err)
	}

	var jwtPayload GoogleJwtPayload

	_, _, err = new(jwtLib.Parser).ParseUnverified(dataResponse.IDToken, &jwtPayload)
	if err != nil {
		return nil, nil, err
	}

	err = jwtPayload.Valid()
	if err != nil {
		return nil, nil, err
	}

	return &jwtPayload, nil, nil
}
