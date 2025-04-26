package factory_model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/factory/factory_model"
	"github.com/stretchr/testify/require"
)

func TestCardShouldGenerate(t *testing.T) {
	card := factory_model.NewCard("4111111111111111", "9", "25", "123", "Test Holder", "12345678909")
	require.NotNil(t, card)
	require.Equal(t, "4111111111111111", card.Number)
	require.Equal(t, "09", card.ExpirationMonth)
	require.Equal(t, "2025", card.ExpirationYear)
	require.Equal(t, "123", card.CVV)
	require.Equal(t, "Test Holder", card.Holder.Name)
	require.Equal(t, "12345678909", card.Holder.Document)
}
