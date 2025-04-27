package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationMustUnmarshal(t *testing.T) {
	authentication := response.NewAuthentication()
	require.NoError(t, authentication.Unmarshal([]byte(`{"access_token":"token","token_type":"bearer","expires_in":3600}`)))
	require.Equal(t, "token", authentication.AccessToken)
	require.Equal(t, "bearer", authentication.TokenType)
	require.Equal(t, 3600, authentication.ExpiresIn)
}
