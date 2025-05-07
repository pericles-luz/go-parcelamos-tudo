package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrPlanIDIsRequired     = errors.New("plan_id is required")
	ErrChargeTypeIsRequired = errors.New("charge_type is required")
	ErrCardIDIsRequired     = errors.New("card_id is required")
	ErrStartDateIsRequired  = errors.New("start_date is required")
	ErrCustomerIsRequired   = errors.New("customer is required")
)

type Subscription struct {
	PlanID              string    `json:"id_card"`
	CardID              string    `json:"id_plan"`
	ChargeType          string    `json:"charge_type"`
	ExternalReferenceID string    `json:"external_reference_id"`
	StartDate           string    `json:"start_date"`
	Cycles              int       `json:"cycles"`
	Customer            *Customer `json:"customer"`
}

func NewSubscription() *Subscription {
	return &Subscription{}
}

func (s *Subscription) SetPlanID(planID string) *Subscription {
	s.PlanID = planID
	return s
}

func (s *Subscription) SetCardID(cardID string) *Subscription {
	s.CardID = cardID
	return s
}

func (s *Subscription) SetChargeType(chargeType string) *Subscription {
	s.ChargeType = chargeType
	return s
}

func (s *Subscription) SetExternalReferenceID(externalReferenceID string) *Subscription {
	s.ExternalReferenceID = externalReferenceID
	return s
}

func (s *Subscription) SetStartDate(startDate string) *Subscription {
	s.StartDate = startDate
	return s
}

func (s *Subscription) SetCycles(cycles int) *Subscription {
	s.Cycles = cycles
	return s
}

func (s *Subscription) SetCustomer(customer *Customer) *Subscription {
	s.Customer = customer
	return s
}

func (s *Subscription) Validate() error {
	if s.PlanID == "" {
		return ErrPlanIDIsRequired
	}

	if !s.IsCreditCard() && !s.IsPix() {
		return ErrChargeTypeIsRequired
	}

	if s.IsCreditCard() && s.CardID == "" {
		return ErrCardIDIsRequired
	}

	if s.StartDate == "" {
		return ErrStartDateIsRequired
	}
	if s.Customer == nil {
		return ErrCustomerIsRequired
	}
	return nil
}

func (s *Subscription) IsCreditCard() bool {
	return s.ChargeType == "credit_card"
}

func (s *Subscription) IsPix() bool {
	return s.ChargeType == "pix"
}

func (s *Subscription) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"id_plan":               s.PlanID,
		"charge_type":           s.ChargeType,
		"external_reference_id": s.ExternalReferenceID,
		"start_date":            s.StartDate,
		"customer":              s.Customer.ToMap(),
	}
	if s.IsCreditCard() {
		result["id_card"] = s.CardID
	}
	if s.Cycles > 0 {
		result["cycles"] = s.Cycles
	}
	return result
}

func (s *Subscription) BindFromMap(data map[string]interface{}) error {
	if v, ok := data["plan_id"].(string); ok {
		s.PlanID = v
	}
	if v, ok := data["card_id"].(string); ok {
		s.CardID = v
	}
	if v, ok := data["charge_type"].(string); ok {
		s.ChargeType = v
	}
	if v, ok := data["external_reference_id"].(string); ok {
		s.ExternalReferenceID = v
	}
	if v, ok := data["start_date"].(string); ok {
		s.StartDate = v
	}
	if v, ok := data["cycles"].(int); ok {
		s.Cycles = v
	}
	if v, ok := data["customer"].(map[string]interface{}); ok {
		customer := NewCustomer()
		if err := customer.BindFromMap(v); err != nil {
			return err
		}
		s.Customer = customer
	}
	return s.Validate()
}

func (s *Subscription) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return s.Validate()
}
