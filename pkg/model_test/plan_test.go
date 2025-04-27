package model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestPlanMustValidate(t *testing.T) {
	plan := model.NewPlan().SetAmount(100).SetDaysUntilDue(10).
		SetDescription("Test").SetName("Test")
	require.NoError(t, plan.Validate())
}

func TestPlanShouldNotValidateIfAmountIsZero(t *testing.T) {
	plan := model.NewPlan().SetDaysUntilDue(10).
		SetDescription("Test").SetName("Test")
	require.ErrorIs(t, plan.Validate(), model.ErrAmountIsRequired)
}

func TestPlanShouldNotValidateIfNameIsEmpty(t *testing.T) {
	plan := model.NewPlan().SetAmount(100).SetDaysUntilDue(10).
		SetDescription("Test")
	require.ErrorIs(t, plan.Validate(), model.ErrNameIsRequired)
}

func TestPlanShouldNotValidateIfDescriptionIsEmpty(t *testing.T) {
	plan := model.NewPlan().SetAmount(100).SetDaysUntilDue(10).
		SetName("Test")
	require.ErrorIs(t, plan.Validate(), model.ErrDescriptionIsRequired)
}

func TestPlanMustGenerateToMap(t *testing.T) {
	plan := model.NewPlan().SetAmount(100).SetDaysUntilDue(10).
		SetDescription("Test").SetName("Test")
	require.Equal(t, map[string]interface{}{
		"name":                  "Test",
		"description":           "Test",
		"external_reference_id": "",
		"currency":              "BRL",
		"amount":                uint32(100),
		"period":                "monthly",
		"days_until_due":        uint32(10),
	}, plan.ToMap())
}
