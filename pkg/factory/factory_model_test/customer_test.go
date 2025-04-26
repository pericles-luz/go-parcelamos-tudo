package factory_model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/factory/factory_model"
	"github.com/stretchr/testify/require"
)

func TestCustomerShouldCreateSlim(t *testing.T) {
	customer := factory_model.NewSlimCustomer("Test Customer", "12345678909")
	require.NotNil(t, customer)
	require.Equal(t, "Test Customer", customer.Name)
	require.Equal(t, "12345678909", customer.Document)
	require.Empty(t, customer.Email)
	require.Empty(t, customer.IP)
	require.NoError(t, customer.Validate())
}
