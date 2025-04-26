package model

import "errors"

var (
	ErrDescriptionIsRequired  = errors.New("description is required")
	ErrCurrencyIsRequired     = errors.New("currency is required")
	ErrAmountIsRequired       = errors.New("amount is required")
	ErrPeriodIsRequired       = errors.New("period is required")
	ErrDaysUntilDueIsRequired = errors.New("days_until_due is required")
)

type Plan struct {
	ID                  string `json:"id_plan"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	ExternalReferenceID string `json:"external_reference_id"`
	Currency            string `json:"currency"`
	Period              string `json:"period"`
	Amount              uint32 `json:"amount"`
	DaysUntilDue        uint8  `json:"days_until_due"`
}

func NewPlan() *Plan {
	return &Plan{
		Currency: "BRL",
		Period:   "monthly",
	}
}

func (p *Plan) SetName(name string) *Plan {
	p.Name = name
	return p
}

func (p *Plan) SetDescription(description string) *Plan {
	p.Description = description
	return p
}

func (p *Plan) SetExternalReferenceID(externalReferenceID string) *Plan {
	p.ExternalReferenceID = externalReferenceID
	return p
}

func (p *Plan) SetCurrency(currency string) *Plan {
	p.Currency = currency
	return p
}

func (p *Plan) SetAmount(amount uint32) *Plan {
	p.Amount = amount
	return p
}

func (p *Plan) SetPeriod(period string) *Plan {
	p.Period = period
	return p
}

func (p *Plan) SetDaysUntilDue(daysUntilDue uint8) *Plan {
	p.DaysUntilDue = daysUntilDue
	return p
}

func (p *Plan) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":                  p.Name,
		"description":           p.Description,
		"external_reference_id": p.ExternalReferenceID,
		"currency":              p.Currency,
		"amount":                p.Amount,
		"period":                p.Period,
		"days_until_due":        p.DaysUntilDue,
	}
}

func (p *Plan) IsMonthly() bool {
	return p.Period == "monthly"
}

func (p *Plan) IsYearly() bool {
	return p.Period == "yearly"
}

func (p *Plan) Validate() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Description == "" {
		return ErrDescriptionIsRequired
	}
	if p.Currency != "BRL" {
		return ErrCurrencyIsRequired
	}
	if p.Amount == 0 {
		return ErrAmountIsRequired
	}
	if !p.IsMonthly() && !p.IsYearly() {
		return ErrPeriodIsRequired
	}
	return nil
}
