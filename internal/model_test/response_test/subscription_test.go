package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model/response"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionMustValidate(t *testing.T) {
	subscription := response.NewSubscription()
	subscription.Success = true
	subscription.Subscrption.ID = "subscription_id"
	subscription.Subscrption.CyclesDone = 1
	require.NoError(t, subscription.Validate())
}

func TestSubscriptionShouldNotValidateIfSuccessIsFalse(t *testing.T) {
	subscription := response.NewSubscription()
	subscription.Subscrption.ID = "subscription_id"
	subscription.Subscrption.CyclesDone = 1
	require.ErrorIs(t, subscription.Validate(), response.ErrSubscriptionFailed)
}

func TestSubscriptionShouldNotValidateIfSubscriptionIDIsEmpty(t *testing.T) {
	subscription := response.NewSubscription()
	subscription.Success = true
	subscription.Subscrption.CyclesDone = 1
	require.ErrorIs(t, subscription.Validate(), response.ErrSubscriptionIDIsRequired)
}

func TestSubscriptionMustUnmarshal(t *testing.T) {
	subscription := response.NewSubscription()
	require.NoError(t, subscription.Unmarshal([]byte(`{"success":true,"subscription":{"id_subscription":"subscription_id","cycles_done":1}}`)))
	require.Equal(t, true, subscription.Success)
	require.Equal(t, "subscription_id", subscription.Subscrption.ID)
	require.Equal(t, 1, subscription.Subscrption.CyclesDone)
}
