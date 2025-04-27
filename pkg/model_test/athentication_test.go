package model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationMustValidate(t *testing.T) {
	auth := model.NewAuthentication().
		SetClientID("client_id").SetClientSecret("client_secret").
		AddScope("scope1").AddScope("scope2")
	require.NoError(t, auth.Validate())
}

func TestAuthenticationShouldNotValidateIfGrantTypeIsEmpty(t *testing.T) {
	auth := model.NewAuthentication().
		SetClientID("client_id").SetClientSecret("client_secret").
		AddScope("scope1").AddScope("scope2").SetGrantType("")
	require.ErrorIs(t, auth.Validate(), model.ErrGrantTypeIsRequired)
}

func TestAuthenticationShouldNotValidateIfClientIDIsEmpty(t *testing.T) {
	auth := model.NewAuthentication().
		SetClientSecret("client_secret").
		AddScope("scope1").AddScope("scope2").SetGrantType("client_credentials")
	require.ErrorIs(t, auth.Validate(), model.ErrClientIDIsRequired)
}

func TestAuthenticationShouldNotValidateIfClientSecretIsEmpty(t *testing.T) {
	auth := model.NewAuthentication().
		SetClientID("client_id").
		AddScope("scope1").AddScope("scope2").SetGrantType("client_credentials")
	require.ErrorIs(t, auth.Validate(), model.ErrClientSecretIsRequired)
}

func TestAuthenticationShouldNotValidateIfScopesAreEmpty(t *testing.T) {
	auth := model.NewAuthentication().
		SetClientID("client_id").SetClientSecret("client_secret").
		SetGrantType("client_credentials")
	require.ErrorIs(t, auth.Validate(), model.ErrScopesIsRequired)
}

func TestAuthenticationMustUnmarshal(t *testing.T) {
	auth := model.NewAuthentication()
	require.NoError(t, auth.Unmarshal([]byte(`{"grant_type":"client_credentials","client_id":"client_id","client_secret":"client_secret","scopes":["scope1","scope2"]}`)))
	require.Equal(t, "client_credentials", auth.GrantType)
	require.Equal(t, "client_id", auth.ClientID)
	require.Equal(t, "client_secret", auth.ClientSecret)
	require.Equal(t, []string{"scope1", "scope2"}, auth.Scopes)
}
