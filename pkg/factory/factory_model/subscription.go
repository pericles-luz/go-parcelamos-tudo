package factory_model

import (
	"time"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
)

func NewSubscriptionWithCard(planID, cardID, customerName, customerDocument string) *model.Subscription {
	customer := NewSlimCustomer(customerName, customerDocument)
	subscription := model.NewSubscription()
	subscription.SetPlanID(planID).
		SetCardID(cardID).
		SetChargeType("credit_card").
		SetStartDate(time.Now().Format("2006-01-02")).
		SetCycles(0).
		SetCustomer(customer)
	if err := subscription.Validate(); err != nil {
		return nil
	}
	return subscription
}

func NewSubscriptionWithPix(planID, customerName, customerDocument, customerEmail string) *model.Subscription {
	customer := NewCustomer(customerName, customerDocument, customerEmail, "")
	subscription := model.NewSubscription()
	subscription.SetPlanID(planID).
		SetChargeType("pix").
		SetStartDate(time.Now().Format("2006-01-02")).
		SetCycles(0).
		SetCustomer(customer)
	if err := subscription.Validate(); err != nil {
		return nil
	}
	return subscription
}
