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
