package rest_test

import (
	"testing"
	"time"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/pericles-luz/go-parcelamos-tudo/pkg/rest"
	"github.com/stretchr/testify/require"
)

func TestTokenMustCreateFromAuthenticationReturn(t *testing.T) {
	auth := &response.Authentication{
		AccessToken: "1234567890",
		ExpiresIn:   3600,
	}
	token := rest.NewToken(auth)
	require.NotNil(t, token)
	require.Equal(t, "1234567890", token.GetKey())
	require.True(t, token.IsValid())
	require.Equal(t, time.Now().UTC().Add(time.Duration(3600)*time.Second).Format("2006-01-02 15:04:05"), token.GetValidity())
}

func TestTokenShouldNotValidateIfKeyLengthIsZero(t *testing.T) {
	token := &rest.Token{}
	token.SetKey("")
	token.SetExpiresIn(3600)
	require.False(t, token.IsValid())
}

func TestTokenShouldNotValidateIfValidityIsZero(t *testing.T) {
	token := &rest.Token{}
	token.SetKey("1234567890")
	require.False(t, token.IsValid())
}

func TestTokenShouldNotValidateIfValidityIsInThePast(t *testing.T) {
	token := &rest.Token{}
	token.SetKey("1234567890")
	token.SetExpiresIn(-1)
	require.False(t, token.IsValid())
}
