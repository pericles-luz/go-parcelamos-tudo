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

func TestWebhookShouldUnmarshalRealInvoiceOpened(t *testing.T) {
	data := []byte(`{
	"context":"invoice",
	"event":"invoice.opened",
	"establishment":"453a685b-1392-4084-ae78-41ae144b3bc8",
	"timestamp":"2025-06-02T13:37:03.114Z",
	"version":"1",
	"data":{
		"id_invoice":"inv_2xxEeqsyEQymUsIv9TdyWQA9Qmk",
		"id_plan":"pln_2wNMJaot9NT0rfbpSFqYsc1vDgm",
		"id_subscription":"sub_2xxEet1aJBv9O6qUffSDPlRxSkT",
		"id_card":null,
		"status":"open",
		"charge_type":"pix",
		"reference_date":"2025-06-02T00:00:00Z",
		"charge_date":null,
		"charge_intent":0,
		"amount":100,
		"currency":"BRL",
		"created_at":"2025-06-02T13:37:02.879884Z",
		"updated_at":"2025-06-02T13:37:02.907125995Z",
		"establishment_id":"453a685b-1392-4084-ae78-41ae144b3bc8"
	}
}`)

	entity := webhook.NewWebhookInvoice()
	require.NoError(t, entity.Unmarshal(data), "should unmarshal webhook data")

	require.Equal(t, "invoice", entity.Context)
	require.Equal(t, "invoice.opened", entity.Event)
	require.Equal(t, "453a685b-1392-4084-ae78-41ae144b3bc8", entity.Establishment)
	require.Equal(t, "2025-06-02T13:37:03.114Z", entity.Timestamp)
	require.Equal(t, "1", entity.Version)

	invoice := entity.Data
	require.Equal(t, "inv_2xxEeqsyEQymUsIv9TdyWQA9Qmk", invoice.ID)
	require.Equal(t, "inv_2xxEeqsyEQymUsIv9TdyWQA9Qmk", entity.ID())
	require.Equal(t, "pln_2wNMJaot9NT0rfbpSFqYsc1vDgm", invoice.PlanID)
	require.Equal(t, "sub_2xxEet1aJBv9O6qUffSDPlRxSkT", invoice.SubscriptionID)
	require.Equal(t, "", invoice.CardID)
	require.Equal(t, "open", invoice.Status)
	require.Equal(t, "2025-06-02T00:00:00Z", invoice.ReferenceDate)
	require.Equal(t, "2025-06-02T00:00:00Z", entity.PaymentDate())
	require.Equal(t, 0, invoice.ChargeIntent)
	require.Equal(t, 100, invoice.Amount)
	require.Equal(t, 100, entity.Amount())
	require.Equal(t, "BRL", invoice.Currency)
	require.Equal(t, "2025-06-02T13:37:02.879884Z", invoice.CreatedAt)
	require.Equal(t, "2025-06-02T13:37:02.907125995Z", invoice.UpdatedAt)
}

func TestWebhookShouldUnmarshalRealInvoiceChargedPix(t *testing.T) {
	data := []byte(`{
	"context":"invoice",
	"event":"invoice.charge.pix",
	"establishment":"453a685b-1392-4084-ae78-41ae144b3bc8",
	"timestamp":"2025-06-11T13:55:52.447Z",
	"version":"1",
	"data":{
		"id_invoice":"inv_2yMh3oFmD6CLnt57ctZKqmU6kXN",
		"id_plan":"pln_2wNMJaot9NT0rfbpSFqYsc1vDgm",
		"id_subscription":"sub_2yMh3j0o9WPi1NgY7rwrBNaHRsl",
		"id_card":null,
		"status":"open",
		"charge_type":"pix",
		"reference_date":"2025-06-11T00:00:00Z",
		"charge_date":null,
		"charge_intent":0,
		"amount":100,
		"currency":"BRL",
		"created_at":"2025-06-11T13:55:52.084545Z",
		"updated_at":"2025-06-11T13:55:52.116980234Z",
		"establishment_id":"453a685b-1392-4084-ae78-41ae144b3bc8"
	}
}`)

	entity := webhook.NewWebhookInvoice()
	require.NoError(t, entity.Unmarshal(data), "should unmarshal webhook data")

	require.Equal(t, "invoice", entity.Context)
	require.Equal(t, "invoice.charge.pix", entity.Event)
	require.Equal(t, "453a685b-1392-4084-ae78-41ae144b3bc8", entity.Establishment)
	require.Equal(t, "2025-06-11T13:55:52.447Z", entity.Timestamp)
	require.Equal(t, "1", entity.Version)

	invoice := entity.Data
	require.Equal(t, "inv_2yMh3oFmD6CLnt57ctZKqmU6kXN", invoice.ID)
	require.Equal(t, "pln_2wNMJaot9NT0rfbpSFqYsc1vDgm", invoice.PlanID)
	require.Equal(t, "sub_2yMh3j0o9WPi1NgY7rwrBNaHRsl", invoice.SubscriptionID)
	require.Equal(t, "", invoice.CardID)
	require.Equal(t, "open", invoice.Status)
	require.Equal(t, "2025-06-11T00:00:00Z", invoice.ReferenceDate)
	require.Equal(t, 0, invoice.ChargeIntent)
	require.Equal(t, 100, invoice.Amount)
	require.Equal(t, "BRL", invoice.Currency)
	require.Equal(t, "2025-06-11T13:55:52.084545Z", invoice.CreatedAt)
	require.Equal(t, "2025-06-11T13:55:52.116980234Z", invoice.UpdatedAt)
}
