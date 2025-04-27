package response

import (
	"encoding/json"
	"errors"
)

var (
	ErrSuccessIsRequired = errors.New("success is required")
	ErrIDIsRequired      = errors.New("id is required")
)

type Plan struct {
	Success    bool   `json:"success"`
	ID         string `json:"id_plan"`
	MerchantID string `json:"id_merchant"`
}

func NewPlan() *Plan {
	return &Plan{}
}

func (p *Plan) Validate() error {
	if !p.Success {
		return ErrSuccessIsRequired
	}
	if p.ID == "" {
		return ErrIDIsRequired
	}
	return nil
}

func (p *Plan) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return p.Validate()
}
