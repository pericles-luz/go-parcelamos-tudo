package response

import (
	"encoding/json"
	"errors"
)

var (
	ErrAccessTokenIsRequired = errors.New("access_token is required")
	ErrTokenTypeIsRequired   = errors.New("token_type is required")
	ErrExpiresInIsRequired   = errors.New("expires_in is required")
)

type Authentication struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func (a *Authentication) Unmarshal(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a *Authentication) Validate() error {
	if a.AccessToken == "" {
		return ErrAccessTokenIsRequired
	}
	if a.TokenType == "" {
		return ErrTokenTypeIsRequired
	}
	if a.ExpiresIn == 0 {
		return ErrExpiresInIsRequired
	}
	return nil
}
