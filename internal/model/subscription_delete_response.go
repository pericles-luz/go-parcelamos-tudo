package model

import "encoding/json"

type SubscriptionDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewSubscriptionDeleteResponse() *SubscriptionDeleteResponse {
	return &SubscriptionDeleteResponse{}
}

func (s *SubscriptionDeleteResponse) SetSuccess(success bool) *SubscriptionDeleteResponse {
	s.Success = success
	return s
}

func (s *SubscriptionDeleteResponse) SetMessage(message string) *SubscriptionDeleteResponse {
	s.Message = message
	return s
}

func (s *SubscriptionDeleteResponse) Validate() error {
	if !s.Success {
		return ErrSubscriptionFailed
	}
	return nil
}

func (s *SubscriptionDeleteResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return s.Validate()
}
