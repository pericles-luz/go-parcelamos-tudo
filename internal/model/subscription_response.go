package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrSubscriptionFailed       = errors.New("subscription failed")
	ErrSubscriptionIDIsRequired = errors.New("subscription id is required")
)

type SubscriptionResponse struct {
	Success     bool `json:"success"`
	Subscrption struct {
		ID         string `json:"id_subscription"`
		CyclesDone int    `json:"cycles_done"`
	} `json:"subscription"`
}

func NewSubscriptionResponse() *SubscriptionResponse {
	return &SubscriptionResponse{}
}

func (s *SubscriptionResponse) SetSuccess(success bool) *SubscriptionResponse {
	s.Success = success
	return s
}

func (s *SubscriptionResponse) SetSubscriptionID(subscriptionID string) *SubscriptionResponse {
	s.Subscrption.ID = subscriptionID
	return s
}

func (s *SubscriptionResponse) SetCyclesDone(cyclesDone int) *SubscriptionResponse {
	s.Subscrption.CyclesDone = cyclesDone
	return s
}

func (s *SubscriptionResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return s.Validate()
}

func (s *SubscriptionResponse) Validate() error {
	if !s.Success {
		return ErrSubscriptionFailed
	}
	if s.Subscrption.ID == "" {
		return ErrSubscriptionIDIsRequired
	}
	return nil
}
