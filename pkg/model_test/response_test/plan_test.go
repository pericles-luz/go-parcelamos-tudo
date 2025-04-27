package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/stretchr/testify/require"
)

func TestPlanMustValidate(t *testing.T) {
	plan := response.NewPlan()
	plan.ID = "1"
	plan.Success = true
	require.NoError(t, plan.Validate())
}

func TestPlanShouldNotValidateIfIDIsEmpty(t *testing.T) {
	plan := response.NewPlan()
	plan.ID = ""
	plan.Success = true
	require.ErrorIs(t, plan.Validate(), response.ErrIDIsRequired)
}

func TestPlanMustUnmarshal(t *testing.T) {
	plan := response.NewPlan()
	require.NoError(t, plan.Unmarshal([]byte(`{"id_plan":"1","success":true}`)))
	require.Equal(t, "1", plan.ID)
	require.True(t, plan.Success)
}
