package factory_model

import "github.com/pericles-luz/go-parcelamos-tudo/pkg/model"

func NewCard(number, expirationMonth, expirationYear, cvv, holderName, holderDocument string) *model.Card {
	result := model.NewCard()
	if len(expirationYear) == 2 {
		expirationYear = "20" + expirationYear
	}
	if len(expirationMonth) == 1 {
		expirationMonth = "0" + expirationMonth
	}
	result.SetNumber(number).
		SetExpirationMonth(expirationMonth).
		SetExpirationYear(expirationYear).
		SetCVV(cvv).
		SetHolderName(holderName).
		SetHolderDocument(holderDocument)
	if err := result.Validate(); err != nil {
		return nil
	}
	return result
}
