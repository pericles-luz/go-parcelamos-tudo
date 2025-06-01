package webhook_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/webhook"
	"github.com/stretchr/testify/require"
)

func TestWebhookShouldUnmarshal(t *testing.T) {
	data := []byte(`{
		"context": "test_context",
		"event": "test_event",
		"establishment": "test_establishment",
		"timestamp": "2023-10-01T00:00:00Z",
		"version": "1.0",
		"data": {
			"id_invoice": "12345",
			"id_plan": "67890",
			"id_subscription": "112233",
			"id_card": "445566",
			"status": "paid",
			"reference_date": "2023-10-01",
			"charge_intent": 1,
			"amount": 1000,
			"currency": "USD",
			"created_at": "2023-10-01T00:00:00Z",
			"updated_at": "2023-10-01T00:00:00Z"
		}
	}`)

	entity := webhook.NewWebhookInvoice()
	require.NoError(t, entity.Unmarshal(data), "should unmarshal webhook data")

	require.Equal(t, "test_context", entity.Context)
	require.Equal(t, "test_event", entity.Event)
	require.Equal(t, "test_establishment", entity.Establishment)
	require.Equal(t, "2023-10-01T00:00:00Z", entity.Timestamp)
	require.Equal(t, "1.0", entity.Version)

	invoice := entity.Data
	require.Equal(t, "12345", invoice.ID)
	require.Equal(t, "67890", invoice.PlanID)
	require.Equal(t, "112233", invoice.SubscriptionID)
	require.Equal(t, "445566", invoice.CardID)
	require.Equal(t, "paid", invoice.Status)
	require.Equal(t, "2023-10-01", invoice.ReferenceDate)
	require.Equal(t, 1, invoice.ChargeIntent)
	require.Equal(t, 1000, invoice.Amount)
	require.Equal(t, "USD", invoice.Currency)
	require.Equal(t, "2023-10-01T00:00:00Z", invoice.CreatedAt)
	require.Equal(t, "2023-10-01T00:00:00Z", invoice.UpdatedAt)
}
