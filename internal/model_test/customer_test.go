package modeltest

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/stretchr/testify/require"
)

func TestCustomerMustValidate(t *testing.T) {
	customer := model.NewCustomer().
		SetName("name").SetEmail("email").
		SetDocument("document").SetIP("ip")
	require.NoError(t, customer.Validate())
}

func TestCustomerShouldNotValidateIfNameIsEmpty(t *testing.T) {
	customer := model.NewCustomer().
		SetEmail("email").SetDocument("document").SetIP("ip")
	require.ErrorIs(t, customer.Validate(), model.ErrNameIsRequired)
}

func TestCustomerShouldNotValidateIfEmailIsEmpty(t *testing.T) {
	customer := model.NewCustomer().
		SetName("name").SetDocument("document").SetIP("ip")
	require.ErrorIs(t, customer.Validate(), model.ErrEmailIsRequired)
}

func TestCustomerShouldNotValidateIfDocumentIsEmpty(t *testing.T) {
	customer := model.NewCustomer().
		SetName("name").SetEmail("email").SetIP("ip")
	require.ErrorIs(t, customer.Validate(), model.ErrDocumentIsRequired)
}

func TestCustomerShouldNotValidateIfIPIsEmpty(t *testing.T) {
	customer := model.NewCustomer().
		SetName("name").SetEmail("email").SetDocument("document")
	require.ErrorIs(t, customer.Validate(), model.ErrIPIsRequired)
}

func TestCustomerMustBindFromMap(t *testing.T) {
	customer := model.NewCustomer()
	require.NoError(t, customer.BindFromMap(map[string]interface{}{
		"name":     "name",
		"email":    "email",
		"document": "document",
		"ip":       "ip",
	}))
	require.Equal(t, "name", customer.Name)
	require.Equal(t, "email", customer.Email)
	require.Equal(t, "document", customer.Document)
	require.Equal(t, "ip", customer.IP)
}
