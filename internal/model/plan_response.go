package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrSuccessIsRequired = errors.New("success is required")
	ErrIDIsRequired      = errors.New("id is required")
)

type PlanResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id_plan"`
}

func NewPlanResponse() *PlanResponse {
	return &PlanResponse{}
}

func (p *PlanResponse) SetSuccess(success bool) *PlanResponse {
	p.Success = success
	return p
}

func (p *PlanResponse) SetID(id string) *PlanResponse {
	p.ID = id
	return p
}

func (p *PlanResponse) Validate() error {
	if !p.Success {
		return ErrSuccessIsRequired
	}
	if p.ID == "" {
		return ErrIDIsRequired
	}
	return nil
}

func (p *PlanResponse) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return p.Validate()
}
