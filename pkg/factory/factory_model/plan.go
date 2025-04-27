package factory_model

import "github.com/pericles-luz/go-parcelamos-tudo/pkg/model"

func NewMonthlyPlan(name, description, externalReferenceID string, amount uint32) *model.Plan {
	plan := model.NewPlan()
	plan.SetName(name).
		SetDescription(description).
		SetExternalReferenceID(externalReferenceID).
		SetAmount(amount).
		SetCurrency("BRL").
		SetPeriod("monthly").
		SetDaysUntilDue(1)
	if err := plan.Validate(); err != nil {
		return nil
	}
	return plan
}
