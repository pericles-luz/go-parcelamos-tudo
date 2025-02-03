package modeltest

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionResponseMustValidate(t *testing.T) {
	subscriptionReturn := model.NewSubscriptionResponse().
		SetSuccess(true).
		SetSubscriptionID("subscription_id").
		SetCyclesDone(1)
	require.NoError(t, subscriptionReturn.Validate())
}

func TestSubscriptionResponseShouldNotValidateIfSuccessIsFalse(t *testing.T) {
	subscriptionReturn := model.NewSubscriptionResponse().
		SetSubscriptionID("subscription_id").
		SetCyclesDone(1)
	require.ErrorIs(t, subscriptionReturn.Validate(), model.ErrSubscriptionFailed)
}

func TestSubscriptionResponseShouldNotValidateIfSubscriptionIDIsEmpty(t *testing.T) {
	subscriptionReturn := model.NewSubscriptionResponse().
		SetSuccess(true).
		SetCyclesDone(1)
	require.ErrorIs(t, subscriptionReturn.Validate(), model.ErrSubscriptionIDIsRequired)
}

func TestSubscriptionResponseMustUnmarshal(t *testing.T) {
	subscriptionReturn := model.NewSubscriptionResponse()
	require.NoError(t, subscriptionReturn.Unmarshal([]byte(`{"success":true,"subscription":{"id_subscription":"subscription_id","cycles_done":1}}`)))
	require.Equal(t, true, subscriptionReturn.Success)
	require.Equal(t, "subscription_id", subscriptionReturn.Subscrption.ID)
	require.Equal(t, 1, subscriptionReturn.Subscrption.CyclesDone)
}
