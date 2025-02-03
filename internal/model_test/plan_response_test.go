package modeltest

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/stretchr/testify/require"
)

func TestPlanResponseMustValidate(t *testing.T) {
	planResponse := model.NewPlanResponse().SetID("1").SetSuccess(true)
	require.NoError(t, planResponse.Validate())
}

func TestPlanResponseShouldNotValidateIfIDIsEmpty(t *testing.T) {
	planResponse := model.NewPlanResponse().SetSuccess(true)
	require.ErrorIs(t, planResponse.Validate(), model.ErrIDIsRequired)
}

func TestPlanResponseMustUnmarshal(t *testing.T) {
	planResponse := model.NewPlanResponse()
	require.NoError(t, planResponse.Unmarshal([]byte(`{"id_plan":"1","success":true}`)))
	require.Equal(t, "1", planResponse.ID)
	require.True(t, planResponse.Success)
}
