package factory_model_test

import (
	"testing"
	"time"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/factory/factory_model"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionShouldCreateWithCard(t *testing.T) {
	subscription := factory_model.NewSubscriptionWithCard(
		"plan id",
		"card id",
		"Test Customer",
		"12345678909",
	)
	require.NotNil(t, subscription)
	require.Equal(t, "plan id", subscription.PlanID)
	require.Equal(t, "card id", subscription.CardID)
	require.Equal(t, "Test Customer", subscription.Customer.Name)
	require.Equal(t, "12345678909", subscription.Customer.Document)
	require.Equal(t, time.Now().Format("2006-01-02"), subscription.StartDate)
	require.Empty(t, subscription.ExternalReferenceID)
	require.Equal(t, 0, subscription.Cycles)
	require.NoError(t, subscription.Validate())
	require.Equal(t, "credit_card", subscription.ChargeType)
	require.Equal(t, "Test Customer", subscription.Customer.Name)
}

func TestSubscriptionShouldCreateWithPix(t *testing.T) {
	subscription := factory_model.NewSubscriptionWithPix(
		"plan id",
		"Test Customer",
		"12345678909",
		"teste@testando.com",
	)
	require.NotNil(t, subscription)
	require.Equal(t, "plan id", subscription.PlanID)
	require.Equal(t, "Test Customer", subscription.Customer.Name)
	require.Equal(t, "12345678909", subscription.Customer.Document)
	require.Equal(t, time.Now().Format("2006-01-02"), subscription.StartDate)
	require.Empty(t, subscription.ExternalReferenceID)
	require.Equal(t, 0, subscription.Cycles)
	require.NoError(t, subscription.Validate())
	require.Equal(t, "pix", subscription.ChargeType)
	require.Equal(t, "Test Customer", subscription.Customer.Name)
	require.Equal(t, "12345678909", subscription.Customer.Document)
	require.Equal(t, "teste@testando.com", subscription.Customer.Email)
}
