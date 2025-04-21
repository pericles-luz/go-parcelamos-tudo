package modeltest

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model/response"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionDeleteMustUnmarshal(t *testing.T) {
	subscriptionDelete := response.NewSubscriptionDelete()
	require.NoError(t, subscriptionDelete.Unmarshal([]byte(`{"success":true,"message":"message"}`)))
	require.Equal(t, true, subscriptionDelete.Success)
	require.Equal(t, "message", subscriptionDelete.Message)
}

func TestSubscriptionDeleteMustValidate(t *testing.T) {
	subscriptionDelete := response.NewSubscriptionDelete()
	subscriptionDelete.Success = true
	subscriptionDelete.Message = "message"
	require.NoError(t, subscriptionDelete.Validate())
}

func TestSubscriptionDeleteShouldNotValidateIfSuccessIsFalse(t *testing.T) {
	subscriptionDelete := response.NewSubscriptionDelete()
	subscriptionDelete.Message = "message"
	require.ErrorIs(t, subscriptionDelete.Validate(), response.ErrSubscriptionFailed)
}
