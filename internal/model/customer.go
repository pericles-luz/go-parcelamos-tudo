package model

import "errors"

var (
	ErrNameIsRequired     = errors.New("name is required")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrDocumentIsRequired = errors.New("document is required")
	ErrIPIsRequired       = errors.New("ip is required")
)

type Customer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	IP       string `json:"ip"`
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c *Customer) SetName(name string) *Customer {
	c.Name = name
	return c
}

func (c *Customer) SetEmail(email string) *Customer {
	c.Email = email
	return c
}

func (c *Customer) SetDocument(document string) *Customer {
	c.Document = document
	return c
}

func (c *Customer) SetIP(ip string) *Customer {
	c.IP = ip
	return c
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return ErrNameIsRequired
	}
	if c.Email == "" {
		return ErrEmailIsRequired
	}
	if c.Document == "" {
		return ErrDocumentIsRequired
	}
	if c.IP == "" {
		return ErrIPIsRequired
	}
	return nil
}

func (c *Customer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":     c.Name,
		"email":    c.Email,
		"document": c.Document,
		"ip":       c.IP,
	}
}

func (c *Customer) BindFromMap(data map[string]interface{}) error {
	if v, ok := data["name"].(string); ok {
		c.Name = v
	}
	if v, ok := data["email"].(string); ok {
		c.Email = v
	}
	if v, ok := data["document"].(string); ok {
		c.Document = v
	}
	if v, ok := data["ip"].(string); ok {
		c.IP = v
	}
	return c.Validate()
}
