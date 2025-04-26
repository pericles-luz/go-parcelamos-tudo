package factory_model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/factory/factory_model"
	"github.com/stretchr/testify/require"
)

func TestPlanShouldCreateMonthly(t *testing.T) {
	plan := factory_model.NewMonthlyPlan("Test Plan", "Test Description", "Test External Reference ID", 100)
	require.NotNil(t, plan)
	require.Equal(t, "Test Plan", plan.Name)
	require.Equal(t, "Test Description", plan.Description)
	require.Equal(t, "Test External Reference ID", plan.ExternalReferenceID)
	require.Equal(t, "BRL", plan.Currency)
	require.Equal(t, uint32(100), plan.Amount)
	require.Equal(t, "monthly", plan.Period)
	require.Equal(t, uint8(1), plan.DaysUntilDue)
}
