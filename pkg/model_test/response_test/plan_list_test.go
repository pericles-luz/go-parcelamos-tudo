package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/stretchr/testify/require"
)

func TestPlanListMustUnmarshal(t *testing.T) {
	planList := response.NewPlanList()
	require.NoError(t, planList.Unmarshal([]byte(`{"total":1,"page":1,"page_size":1,"data":[{"id_plan":"1","name":"plan_name","description":"plan_description","external_reference_id":"external_id","currency":"BRL","amount":100,"period":"monthly","days_until_due":15}]}`)))
	require.Len(t, planList.Data, 1)
	require.Equal(t, "1", planList.Data[0].ID)
	require.Equal(t, "plan_name", planList.Data[0].Name)
	require.Equal(t, "plan_description", planList.Data[0].Description)
	require.Equal(t, "external_id", planList.Data[0].ExternalReferenceID)
	require.Equal(t, "BRL", planList.Data[0].Currency)
	require.Equal(t, uint32(100), planList.Data[0].Amount)
	require.Equal(t, "monthly", planList.Data[0].Period)
	require.Equal(t, uint32(15), planList.Data[0].DaysUntilDue)
	require.Equal(t, "monthly", planList.Data[0].Period)
	require.Equal(t, 1, planList.Total)
	require.Equal(t, 1, planList.Page)
	require.Equal(t, 1, planList.PageSize)
}
