package model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionMustValidate(t *testing.T) {
	subscription := model.NewSubscription().
		SetPlanID("plan_id").SetCardID("card_id").
		SetChargeType("credit_card").SetStartDate("start_date").
		SetCustomer(model.NewCustomer().SetName("name").SetEmail("email").SetDocument("document").SetIP("ip"))
	require.NoError(t, subscription.Validate())
}

func TestSubscriptionShouldNotValidateIfPlanIDIsEmpty(t *testing.T) {
	subscription := model.NewSubscription().
		SetCardID("card_id").SetChargeType("credit_card").
		SetStartDate("start_date").SetCustomer(model.NewCustomer().SetName("name").SetEmail("email").SetDocument("document").SetIP("ip"))
	require.ErrorIs(t, subscription.Validate(), model.ErrPlanIDIsRequired)
}

func TestSubscriptionShouldNotValidateIfChargeTypeIsCreditCardAndCardIDIsEmpty(t *testing.T) {
	subscription := model.NewSubscription().
		SetPlanID("plan_id").SetChargeType("credit_card").
		SetStartDate("start_date").SetCustomer(model.NewCustomer().SetName("name").SetEmail("email").SetDocument("document").SetIP("ip"))
	require.ErrorIs(t, subscription.Validate(), model.ErrCardIDIsRequired)
}

func TestSubscriptionShouldNotValidateIfChargeTypeIsEmpty(t *testing.T) {
	subscription := model.NewSubscription().
		SetPlanID("plan_id").SetCardID("card_id").
		SetStartDate("start_date").SetCustomer(model.NewCustomer().SetName("name").SetEmail("email").SetDocument("document").SetIP("ip"))
	require.ErrorIs(t, subscription.Validate(), model.ErrChargeTypeIsRequired)
}

func TestSubscriptionShouldNotValidateIfStartDateIsEmpty(t *testing.T) {
	subscription := model.NewSubscription().
		SetPlanID("plan_id").SetCardID("card_id").
		SetChargeType("credit_card").SetCustomer(model.NewCustomer().SetName("name").SetEmail("email").SetDocument("document").SetIP("ip"))
	require.ErrorIs(t, subscription.Validate(), model.ErrStartDateIsRequired)
}

func TestSubscriptionShouldNotValidateIfCustomerIsEmpty(t *testing.T) {
	subscription := model.NewSubscription().
		SetPlanID("plan_id").SetCardID("card_id").
		SetChargeType("credit_card").SetStartDate("start_date")
	require.ErrorIs(t, subscription.Validate(), model.ErrCustomerIsRequired)
}

func TestSubscriptionMustBindFromMap(t *testing.T) {
	subscription := model.NewSubscription()
	require.NoError(t, subscription.BindFromMap(map[string]interface{}{
		"plan_id": "plan_id", "card_id": "card_id",
		"charge_type": "credit_card", "start_date": "start_date",
		"customer": map[string]interface{}{
			"name": "name", "email": "email",
			"document": "document", "ip": "ip",
		},
	}))
	require.Equal(t, "plan_id", subscription.PlanID)
	require.Equal(t, "card_id", subscription.CardID)
	require.Equal(t, "credit_card", subscription.ChargeType)
	require.Equal(t, "start_date", subscription.StartDate)
	require.Equal(t, "name", subscription.Customer.Name)
	require.Equal(t, "email", subscription.Customer.Email)
	require.Equal(t, "document", subscription.Customer.Document)
	require.Equal(t, "ip", subscription.Customer.IP)
}
