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
	restEntity, err := rest.NewRest(engine, utils.GetBaseDirectory("config")+"/sandbox.json", []string{"subscription.create"})
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
