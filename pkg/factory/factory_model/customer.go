package factory_model

import "github.com/pericles-luz/go-parcelamos-tudo/pkg/model"

func NewSlimCustomer(name, document string) *model.Customer {
	customer := model.NewCustomer()
	customer.SetName(name).
		SetDocument(document)
	if err := customer.Validate(); err != nil {
		return nil
	}
	return customer
}

func NewCustomer(name, document, email, ip string) *model.Customer {
	customer := model.NewCustomer()
	customer.SetName(name).
		SetDocument(document).
		SetEmail(email).
		SetIP(ip)
	if err := customer.Validate(); err != nil {
		return nil
	}
	return customer
}
