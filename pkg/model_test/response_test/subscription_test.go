package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionMustValidate(t *testing.T) {
	subscription := response.NewSubscription()
	subscription.Success = true
	subscription.Subscription.ID = "subscription_id"
	subscription.Subscription.CyclesDone = 1
	require.NoError(t, subscription.Validate())
}

func TestSubscriptionShouldNotValidateIfSuccessIsFalse(t *testing.T) {
	subscription := response.NewSubscription()
	subscription.Subscription.ID = "subscription_id"
	subscription.Subscription.CyclesDone = 1
	require.ErrorIs(t, subscription.Validate(), response.ErrSubscriptionFailed)
}

func TestSubscriptionShouldNotValidateIfSubscriptionIDIsEmpty(t *testing.T) {
	subscription := response.NewSubscription()
	subscription.Success = true
	subscription.Subscription.CyclesDone = 1
	require.ErrorIs(t, subscription.Validate(), response.ErrSubscriptionIDIsRequired)
}

func TestSubscriptionMustUnmarshal(t *testing.T) {
	subscription := response.NewSubscription()
	// require.NoError(t, subscription.Unmarshal([]byte(`{"success":true,"subscription":{"id_subscription":"subscription_id","cycles_done":1}}`)))
	require.NoError(t, subscription.Unmarshal([]byte(`{"success":true,"subscription":{"id_subscription":"sub_2wk97XjFXOiTAoOYNvjgsEHtlVA","id_plan":"pln_2wNMJaot9NT0rfbpSFqYsc1vDgm","id_merchant":"mer_2safJVeZjodLYxLcrb9HDTf3Ref","id_card":"crd_2w6QqYlyqkdOgs4BzcmLb9GiG92","status":"active","charge_type":"credit_card","charge_day":7,"external_reference_id":"5e7c2b43-532f-49ed-b305-b813f11ef6bd","start_date":"2025-05-07T00:00:00Z","end_date":null,"cycles":null,"cycles_done":0,"total_amount_charges":0,"total_quantity_charges":0,"customer_name":"Test Customer","customer_document":"80205365078","customer_ip":"170.246.46.237","customer_email":null,"product_description":"Pagamento mensal Test","created_at":"2025-05-06T23:35:37.181687Z","updated_at":"2025-05-06T23:35:37.181687Z","establishment_id":"453a685b-1392-4084-ae78-41ae144b3bc8"},"next_invoice":{"id_invoice":"inv_2wk97ZbyBA9qC7DgYzvyWJfYNlT","id_plan":"pln_2wNMJaot9NT0rfbpSFqYsc1vDgm","id_subscription":"sub_2wk97XjFXOiTAoOYNvjgsEHtlVA","id_card":"crd_2w6QqYlyqkdOgs4BzcmLb9GiG92","status":"waiting","charge_type":"credit_card","reference_date":"2025-05-07T00:00:00Z","charge_date":null,"charge_intent":0,"amount":100,"currency":"BRL","created_at":"2025-05-06T23:35:37.199598Z","updated_at":"2025-05-06T23:35:37.199598Z","establishment_id":"453a685b-1392-4084-ae78-41ae144b3bc8"}}`)))
	require.Equal(t, true, subscription.Success)
	require.Equal(t, "sub_2wk97XjFXOiTAoOYNvjgsEHtlVA", subscription.Subscription.ID)
	require.Equal(t, 0, subscription.Subscription.CyclesDone)
}
