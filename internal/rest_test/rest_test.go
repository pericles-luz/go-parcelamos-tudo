package rest_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/pericles-luz/go-parcelamos-tudo/internal/rest"
	"github.com/stretchr/testify/require"
)

func TestRestAuthenticationShouldFailIfHasNoEngine(t *testing.T) {
	restEntity, err := rest.NewRest(nil, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"subscription.create"})
	require.ErrorIs(t, err, rest.ErrMissingEngine)
	require.Nil(t, restEntity)
}

func TestRestAuthenticationShouldWorkWithCredentials(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"card.create", "plan.search"})
	require.NoError(t, err, "Failed to create rest entity")
	require.NoError(t, restEntity.Authenticate(), "Authentication failed")
	require.NoError(t, restEntity.Authenticate(), "Second Authentication failed")
}

func TestRestShouldCreatePlan(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"plan.create"})
	require.NoError(t, err, "Failed to create rest entity")
	plan := model.NewPlan()
	plan.SetName("Test Plan").
		SetDescription("Test Plan Description").
		SetExternalReferenceID(uuid.NewString()).
		SetCurrency("BRL").
		SetAmount(100).
		SetPeriod("monthly").
		SetDaysUntilDue(30)
	require.NoError(t, plan.Validate(), "Plan validation failed")
	t.Log("Plan: ", plan)
	response, err := restEntity.CreatePlan(plan)
	require.NoError(t, err, "Failed to create plan")
	require.True(t, response.Success, "Plan creation failed")
	t.Log("Plan ID: ", response.ID)
}

func TestRestShouldListPlan(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"plan.search"})
	require.NoError(t, err, "Failed to create rest entity")
	planList, err := restEntity.ListPlan(1, 0, "")
	require.NoError(t, err, "Failed to list plan")
	require.True(t, planList.Total > 0, "Plan list is empty")
}

func TestRestShouldGetPlan(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"plan.read"})
	require.NoError(t, err, "Failed to create rest entity")
	readed, err := restEntity.GetPlan("pln_2w5WdzyvPpgFFhgqJSojlMwGZTz")
	require.NoError(t, err, "Failed to get plan")
	require.Equal(t, "pln_2w5WdzyvPpgFFhgqJSojlMwGZTz", readed.ID, "Plan ID is not equal")
}

func TestRestShouldGetPlanByExternalID(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"plan.search"})
	require.NoError(t, err, "Failed to create rest entity")
	readed, err := restEntity.ListPlan(1, 0, "87ce217c-b881-4c0e-b199-d93a88787e73")
	require.NoError(t, err, "Failed to get plan")
	require.Len(t, readed.Data, 1, "Plan list is empty")
	require.Equal(t, "pln_2w5WdzyvPpgFFhgqJSojlMwGZTz", readed.Data[0].ID, "Plan ID is not equal")
}

func TestRestShouldCreateCard(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"card.create"})
	require.NoError(t, err, "Failed to create rest entity")
	card := model.NewCard()
	card.Number = "4761120000000148"
	card.ExpirationMonth = "07"
	card.ExpirationYear = "2026"
	card.CVV = "476"
	card.Holder.Name = "Test Holder"
	card.Holder.Document = "00000000191"
	require.NoError(t, card.Validate(), "Card validation failed")
	response, err := restEntity.CreateCard(card)
	require.NoError(t, err, "Failed to create card")
	require.NotEmpty(t, response.ID, "Card ID is empty")
	t.Log("Card ID: ", response.ID)
	t.Log("Card: ", response)
}

func TestRestShouldGetCard(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"card.read"})
	require.NoError(t, err, "Failed to create rest entity")
	card, err := restEntity.GetCard("crd_2w6QqYlyqkdOgs4BzcmLb9GiG92")
	require.NoError(t, err, "Failed to get card")
	require.Equal(t, "crd_2w6QqYlyqkdOgs4BzcmLb9GiG92", card.ID, "Card ID is not equal")
	t.Log("Card: ", card)
}

func TestRestShouldSubscribe(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"subscription.create"})
	require.NoError(t, err, "Failed to create rest entity")
	subscription := model.NewSubscription()
	subscription.PlanID = "pln_2w5WdzyvPpgFFhgqJSojlMwGZTz"
	subscription.ChargeType = "credit_card"
	subscription.CardID = "crd_2w6QqYlyqkdOgs4BzcmLb9GiG92"
	subscription.ExternalReferenceID = uuid.NewString()
	subscription.StartDate = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	subscription.Cycles = 0
	customer := model.NewCustomer()
	customer.Name = "Test Customer"
	customer.Document = "80205365078"
	subscription.Customer = customer
	require.NoError(t, subscription.Validate(), "Subscription validation failed")
	response, err := restEntity.Subscribe(subscription)
	require.NoError(t, err, "Failed to create subscription")
	require.NotEmpty(t, response.Subscription.ID, "Subscription ID is empty")
	t.Log("Subscription ID: ", response.Subscription.ID)
	t.Log("Subscription: ", response)
	require.Equal(t, subscription.PlanID, response.Subscription.PlanID, "Subscription Plan ID is not equal")
	require.Equal(t, subscription.ChargeType, response.Subscription.ChargeType, "Subscription Charge Type is not equal")
	require.Equal(t, subscription.CardID, response.Subscription.CardID, "Subscription Card ID is not equal")
	require.Equal(t, subscription.ExternalReferenceID, response.Subscription.ExternalReferenceID, "Subscription External Reference ID is not equal")
}
