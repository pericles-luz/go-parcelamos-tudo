package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrCardNumberIsRequired          = errors.New("card number is required")
	ErrCardSecurityCodeIsRequired    = errors.New("card security code is required")
	ErrCardHolderNameIsRequired      = errors.New("card holder name is required")
	ErrCardHolderDocumentIsRequired  = errors.New("card holder document is required")
	ErrCardExpirationMonthIsRequired = errors.New("card expiration month is required")
	ErrCardExpirationYearIsRequired  = errors.New("card expiration year is required")
)

type Holder struct {
	Name     string `json:"holder_name"`
	Document string `json:"holder_document"`
}

type Card struct {
	Number          string `json:"number"`
	CVV             string `json:"security_code"`
	ID              string `json:"id_card"`
	MerchantID      string `json:"id_merchant"`
	Brand           string `json:"brand"`
	FirstDigits     string `json:"first_digits"`
	LastDigits      string `json:"last_digits"`
	ExpirationMonth string `json:"exp_month"`
	ExpirationYear  string `json:"exp_year"`
	Holder          Holder `json:"holder"`
	ValidUntil      string `json:"valid_until"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func NewCard() *Card {
	return &Card{
		Holder: Holder{},
	}
}

func (c *Card) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	return c.Validate()
}

func (c *Card) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"number":          c.Number,
		"security_code":   c.CVV,
		"holder_name":     c.Holder.Name,
		"holder_document": c.Holder.Document,
		"exp_month":       c.ExpirationMonth,
		"exp_year":        c.ExpirationYear,
	}
}

func (c *Card) Validate() error {
	if c.Number == "" {
		return ErrCardNumberIsRequired
	}
	if c.CVV == "" {
		return ErrCardSecurityCodeIsRequired
	}
	if c.Holder.Name == "" {
		return ErrCardHolderNameIsRequired
	}
	if c.Holder.Document == "" {
		return ErrCardHolderDocumentIsRequired
	}
	if c.ExpirationMonth == "" {
		return ErrCardExpirationMonthIsRequired
	}
	if c.ExpirationYear == "" {
		return ErrCardExpirationYearIsRequired
	}
	return nil
}
