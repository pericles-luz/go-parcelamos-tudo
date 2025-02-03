package modeltest

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionDeleteRespnseMustUnmarshal(t *testing.T) {
	subscriptionDeleteResponse := model.NewSubscriptionDeleteResponse()
	require.NoError(t, subscriptionDeleteResponse.Unmarshal([]byte(`{"success":true,"message":"message"}`)))
	require.Equal(t, true, subscriptionDeleteResponse.Success)
	require.Equal(t, "message", subscriptionDeleteResponse.Message)
}

func TestSubscriptionDeleteRespnseMustValidate(t *testing.T) {
	subscriptionDeleteResponse := model.NewSubscriptionDeleteResponse().
		SetSuccess(true).
		SetMessage("message")
	require.NoError(t, subscriptionDeleteResponse.Validate())
}

func TestSubscriptionDeleteRespnseShouldNotValidateIfSuccessIsFalse(t *testing.T) {
	subscriptionDeleteResponse := model.NewSubscriptionDeleteResponse().
		SetMessage("message")
	require.ErrorIs(t, subscriptionDeleteResponse.Validate(), model.ErrSubscriptionFailed)
}
