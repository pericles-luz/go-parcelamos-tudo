package rest_test

import (
	"os"
	"testing"

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
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"subscription.create", "plan.search"})
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
	planList, err := restEntity.ListPlan(1, 0)
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
	readed, err := restEntity.GetPlan("pln_2tGkIfO2nkBjJxAvSWUrAZ3R5X1")
	require.NoError(t, err, "Failed to get plan")
	require.Equal(t, "plid_test", readed.ID, "Plan ID is not equal")
}

func TestRestShouldCreateCard(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"card.create"})
	require.NoError(t, err, "Failed to create rest entity")
	card := model.NewCard()
	card.Number = "4111111111111111"
	card.ExpirationMonth = "12"
	card.ExpirationYear = "2025"
	card.CVV = "123"
	card.Holder.Name = "Test Holder"
	card.Holder.Document = "12345678901"
	require.NoError(t, card.Validate(), "Card validation failed")
	response, err := restEntity.CreateCard(card)
	require.NoError(t, err, "Failed to create card")
	require.NotEmpty(t, response.ID, "Card ID is empty")
}
