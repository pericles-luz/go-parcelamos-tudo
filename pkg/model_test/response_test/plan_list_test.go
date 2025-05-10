package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/stretchr/testify/require"
)

func TestPlanListMustUnmarshal(t *testing.T) {
	planList := response.NewPlanList()
	require.NoError(t, planList.Unmarshal([]byte(`{"data":[{"id_plan":"pln_2wunEZL6uFj8VU2yB30nFVvgrRX","id_merchant":"mer_2safJVeZjodLYxLcrb9HDTf3Ref","active":true,"name":"Test Plan","description":"Test Plan Description","currency":"BRL","external_reference_id":"40e33ab2-05cc-4b2c-8c68-59c6707c1702","amount":1000,"period":"monthly","days_until_due":1,"splitted":false,"created_at":"2025-05-10T18:03:35.103892Z","updated_at":"2025-05-10T18:03:35.103892Z","removed_at":null}],"page":1,"page_size":25,"total":1}`)))
	require.Len(t, planList.Data, 1)
	require.Equal(t, "pln_2wunEZL6uFj8VU2yB30nFVvgrRX", planList.Data[0].ID)
	require.Equal(t, "Test Plan", planList.Data[0].Name)
	require.Equal(t, "Test Plan Description", planList.Data[0].Description)
	require.Equal(t, "40e33ab2-05cc-4b2c-8c68-59c6707c1702", planList.Data[0].ExternalReferenceID)
	require.Equal(t, "BRL", planList.Data[0].Currency)
	require.Equal(t, uint32(1000), planList.Data[0].Amount)
	require.Equal(t, "monthly", planList.Data[0].Period)
	require.Equal(t, uint8(1), planList.Data[0].DaysUntilDue)
	require.Equal(t, "monthly", planList.Data[0].Period)
	require.Equal(t, 1, planList.Total)
	require.Equal(t, 1, planList.Page)
	require.Equal(t, 25, planList.PageSize)
}
