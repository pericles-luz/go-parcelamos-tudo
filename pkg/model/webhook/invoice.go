package webhook

import "encoding/json"

type WebhookInvoice struct {
	Context       string `json:"context"`
	Event         string `json:"event"`
	Establishment string `json:"establishment"`
	Timestamp     string `json:"timestamp"`
	Version       string `json:"version"`
	Data          struct {
		ID             string `json:"id_invoice"`
		PlanID         string `json:"id_plan"`
		SubscriptionID string `json:"id_subscription"`
		CardID         string `json:"id_card"`
		Status         string `json:"status"`
		ReferenceDate  string `json:"reference_date"`
		ChargeIntent   int    `json:"charge_intent"`
		Amount         int    `json:"amount"`
		Currency       string `json:"currency"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
	} `json:"data"`
}

func NewWebhookInvoice() *WebhookInvoice {
	return &WebhookInvoice{}
}

func (w *WebhookInvoice) Unmarshal(data []byte) error {
	return json.Unmarshal(data, w)
}

func (w *WebhookInvoice) IsOpened() bool {
	return w.Event == "invoice.opened"
}

func (w *WebhookInvoice) IsChargingPix() bool {
	return w.Event == "invoice.charge.pix"
}

func (w *WebhookInvoice) IsPaid() bool {
	return w.Event == "invoice.paid"
}

func (w *WebhookInvoice) IsCancelled() bool {
	return w.Event == "invoice.cancelled"
}

func (w *WebhookInvoice) Amount() int {
	return w.Data.Amount
}

func (w *WebhookInvoice) ID() string {
	return w.Data.ID
}

func (w *WebhookInvoice) PaymentDate() string {
	return w.Data.ReferenceDate
}

func (w *WebhookInvoice) SubscriptionID() string {
	return w.Data.SubscriptionID
}
