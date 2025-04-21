package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrDataIsRequired     = errors.New("data is required")
	ErrPageIsRequired     = errors.New("page is required")
	ErrPageSizeIsRequired = errors.New("page_size is required")
	ErrTotalIsRequired    = errors.New("total is required")
)

type PlanList struct {
	Data     []Plan `json:"data"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Total    int    `json:"total"`
}

func NewPlanList() *PlanList {
	return &PlanList{}
}

func (p *PlanList) Validate() error {
	if p.Data == nil {
		return ErrDataIsRequired
	}
	if p.Page < 0 {
		return ErrPageIsRequired
	}
	if p.PageSize <= 0 {
		return ErrPageSizeIsRequired
	}
	if p.Total < 0 {
		return ErrTotalIsRequired
	}
	return nil
}

func (p *PlanList) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return p.Validate()
}
