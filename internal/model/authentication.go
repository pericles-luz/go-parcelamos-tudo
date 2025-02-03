package model

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrGrantTypeIsRequired    = errors.New("grant_type is required")
	ErrClientIDIsRequired     = errors.New("client_id is required")
	ErrClientSecretIsRequired = errors.New("client_secret is required")
	ErrScopesIsRequired       = errors.New("scopes is required")
)

type Authentication struct {
	GrantType    string   `json:"grant_type"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Scopes       []string `json:"scopes"`
}

func NewAuthentication() *Authentication {
	return &Authentication{
		GrantType: "client_credentials",
		Scopes:    []string{},
	}
}

func (a *Authentication) AddScope(scope string) *Authentication {
	a.Scopes = append(a.Scopes, scope)
	return a
}

func (a *Authentication) SetGrantType(grantType string) *Authentication {
	a.GrantType = grantType
	return a
}

func (a *Authentication) SetClientID(clientID string) *Authentication {
	a.ClientID = clientID
	return a
}

func (a *Authentication) SetClientSecret(clientSecret string) *Authentication {
	a.ClientSecret = clientSecret
	return a
}

func (a *Authentication) Validate() error {
	if a.GrantType == "" {
		return ErrGrantTypeIsRequired
	}
	if a.ClientID == "" {
		return ErrClientIDIsRequired
	}
	if a.ClientSecret == "" {
		return ErrClientSecretIsRequired
	}
	if len(a.Scopes) == 0 {
		return ErrScopesIsRequired
	}
	return nil
}

func (a *Authentication) Unmarshal(raw []byte) error {
	if err := json.Unmarshal(raw, a); err != nil {
		return err
	}
	return a.Validate()
}

func (a *Authentication) ToMap() map[string]string {
	return map[string]string{
		"grant_type":    a.GrantType,
		"client_id":     a.ClientID,
		"client_secret": a.ClientSecret,
		"scopes":        strings.Join(a.Scopes, " "),
	}
}
