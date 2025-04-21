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
	Success     bool `json:"success"`
	Subscrption struct {
		ID         string `json:"id_subscription"`
		CyclesDone int    `json:"cycles_done"`
	} `json:"subscription"`
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
	if s.Subscrption.ID == "" {
		return ErrSubscriptionIDIsRequired
	}
	return nil
}
