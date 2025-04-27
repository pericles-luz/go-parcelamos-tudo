package response

import (
	"encoding/json"
	"errors"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
)

var (
	ErrDataIsRequired     = errors.New("data is required")
	ErrPageIsRequired     = errors.New("page is required")
	ErrPageSizeIsRequired = errors.New("page_size is required")
	ErrTotalIsRequired    = errors.New("total is required")
)

type PlanList struct {
	Data     []model.Plan `json:"data"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
	Total    int          `json:"total"`
}

func NewPlanList() *PlanList {
	return &PlanList{}
}

func (p *PlanList) Validate() error {
	if p.Page < 0 {
		return ErrPageIsRequired
	}
	if p.PageSize <= 0 {
		return ErrPageSizeIsRequired
	}
	if p.Total < 0 {
		return ErrTotalIsRequired
	}
	if p.Total == 0 {
		return nil
	}
	if p.Data == nil {
		return ErrDataIsRequired
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
