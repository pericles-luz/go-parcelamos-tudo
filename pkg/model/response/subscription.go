package response

import (
	"encoding/json"
	"errors"
)

var (
	ErrSubscriptionFailed       = errors.New("subscription failed")
	ErrSubscriptionIDIsRequired = errors.New("subscription id is required")
)

type Subscription struct {
	Success      bool `json:"success"`
	Subscription struct {
		ID                   string `json:"id_subscription"`
		CyclesDone           int    `json:"cycles_done"`
		Status               string `json:"status"`
		ExternalReferenceID  string `json:"external_reference_id"`
		PlanID               string `json:"id_plan"`
		CardID               string `json:"id_card"`
		MerchantID           string `json:"id_merchant"`
		ChargeType           string `json:"charge_type"`
		StartDate            string `json:"start_date"`
		TotalAmountCharges   int    `json:"total_amount_charges"`
		TotalQuantityCharges int    `json:"total_quantity_charges"`
		CustomerName         string `json:"customer_name"`
		CustomerDocument     string `json:"customer_document"`
		CreatedAt            string `json:"created_at"`
	} `json:"subscription"`
	NextInvoice struct {
		ID            string `json:"id_invoice"`
		Status        string `json:"status"`
		ReferenceDate string `json:"reference_date"`
		ChargeIntent  int    `json:"charge_intent"`
	} `json:"next_invoice"`
}

func NewSubscription() *Subscription {
	return &Subscription{}
}

func (s *Subscription) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return s.Validate()
}

func (s *Subscription) Validate() error {
	if !s.Success {
		return ErrSubscriptionFailed
	}
	if s.Subscription.ID == "" {
		return ErrSubscriptionIDIsRequired
	}
	return nil
}
