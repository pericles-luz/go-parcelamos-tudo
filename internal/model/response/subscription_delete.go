package response

import "encoding/json"

type SubscriptionDelete struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewSubscriptionDelete() *SubscriptionDelete {
	return &SubscriptionDelete{}
}

func (s *SubscriptionDelete) Validate() error {
	if !s.Success {
		return ErrSubscriptionFailed
	}
	return nil
}

func (s *SubscriptionDelete) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return s.Validate()
}
