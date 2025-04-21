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
	restEntity := rest.NewRest(nil)
	require.Equal(t, rest.ErrMissingEngine, restEntity.Authenticate(&model.Authentication{
		GrantType:    "client_credentials",
		ClientID:     "1234567890",
		ClientSecret: "1234567890",
		Scopes:       []string{"read", "write"},
	}))
}

func TestRestAuthenticationShouldWorkWithCredentials(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	configJSON, err := os.ReadFile(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to read config file")
	autentication := model.NewAuthentication()
	autentication.AddScope("subscription.create")
	autentication.AddScope("subscription.delete")
	require.NoErrorf(t, autentication.Unmarshal(configJSON), "Failed to unmarshal config file: %v", err)
	require.NoError(t, autentication.Validate(), "Failed to validate authentication")
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	rest := rest.NewRest(engine)
	rest.SetBaseLink(autentication.Link)
	require.NoError(t, rest.Authenticate(autentication), "Authentication failed")
	require.NoError(t, rest.Authenticate(autentication), "Authentication failed")
}

func TestRestShouldCreatePlan(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	configJSON, err := os.ReadFile(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to read config file")
	autentication := model.NewAuthentication()
	autentication.AddScope("plan.create")
	require.NoErrorf(t, autentication.Unmarshal(configJSON), "Failed to unmarshal config file: %v", err)
	require.NoError(t, autentication.Validate(), "Failed to validate authentication")
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	rest := rest.NewRest(engine)
	rest.SetBaseLink(autentication.Link)
	require.NoError(t, rest.Authenticate(autentication), "Authentication failed")
	plan := model.NewPlan()
	plan.SetName("Test Plan").
		SetDescription("Test Plan Description").
		SetExternalReferenceID(uuid.NewString()).
		SetCurrency("BRL").
		SetAmount(100).
		SetPeriod("monthly").
		SetDaysUntilDue(30)
	response, err := rest.CreatePlan(plan)
	require.NoError(t, err, "Failed to create plan")
	require.True(t, response.Success, "Plan creation failed")
	t.Log("Plan ID: ", response.ID)
}
