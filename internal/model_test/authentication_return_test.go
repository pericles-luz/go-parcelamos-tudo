package modeltest

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationReturnMustUnmarshal(t *testing.T) {
	authReturn := model.NewAuthenticationReturn()
	require.NoError(t, authReturn.Unmarshal([]byte(`{"access_token":"token","token_type":"bearer","expires_in":3600}`)))
	require.Equal(t, "token", authReturn.AccessToken)
	require.Equal(t, "bearer", authReturn.TokenType)
	require.Equal(t, 3600, authReturn.ExpiresIn)
}
