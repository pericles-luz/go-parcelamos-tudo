package model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/stretchr/testify/require"
)

func TestCardShouldValidate(t *testing.T) {
	card := model.NewCard()
	card.Number = "1234567890123456"
	card.CVV = "123"
	card.ExpirationMonth = "12"
	card.ExpirationYear = "2025"
	card.Holder.Name = "John Doe"
	card.Holder.Document = "12345678901"
	require.NoError(t, card.Validate(), "Card should be valid")
}

func TestCardShouldNotValidateIfNumberIsEmpty(t *testing.T) {
	card := model.NewCard()
	card.CVV = "123"
	card.ExpirationMonth = "12"
	card.ExpirationYear = "2025"
	card.Holder.Name = "John Doe"
	card.Holder.Document = "12345678901"
	require.ErrorIs(t, card.Validate(), model.ErrCardNumberIsRequired, "Card number should be required")
}

func TestCardShouldNotValidateIfCVVIsEmpty(t *testing.T) {
	card := model.NewCard()
	card.Number = "1234567890123456"
	card.ExpirationMonth = "12"
	card.ExpirationYear = "2025"
	card.Holder.Name = "John Doe"
	card.Holder.Document = "12345678901"
	require.ErrorIs(t, card.Validate(), model.ErrCardSecurityCodeIsRequired, "Card security code should be required")
}

func TestCardShouldNotValidateIfExpirationMonthIsEmpty(t *testing.T) {
	card := model.NewCard()
	card.Number = "1234567890123456"
	card.CVV = "123"
	card.ExpirationYear = "2025"
	card.Holder.Name = "John Doe"
	card.Holder.Document = "12345678901"
	require.ErrorIs(t, card.Validate(), model.ErrCardExpirationMonthIsRequired, "Card expiration month should be required")
}

func TestCardShouldNotValidateIfExpirationYearIsEmpty(t *testing.T) {
	card := model.NewCard()
	card.Number = "1234567890123456"
	card.CVV = "123"
	card.ExpirationMonth = "12"
	card.Holder.Name = "John Doe"
	card.Holder.Document = "12345678901"
	require.ErrorIs(t, card.Validate(), model.ErrCardExpirationYearIsRequired, "Card expiration year should be required")
}

func TestCardShouldNotValidateIfHolderNameIsEmpty(t *testing.T) {
	card := model.NewCard()
	card.Number = "1234567890123456"
	card.CVV = "123"
	card.ExpirationMonth = "12"
	card.ExpirationYear = "2025"
	card.Holder.Document = "12345678901"
	require.ErrorIs(t, card.Validate(), model.ErrCardHolderNameIsRequired, "Card holder name should be required")
}

func TestCardShouldNotValidateIfHolderDocumentIsEmpty(t *testing.T) {
	card := model.NewCard()
	card.Number = "1234567890123456"
	card.CVV = "123"
	card.ExpirationMonth = "12"
	card.ExpirationYear = "2025"
	card.Holder.Name = "John Doe"
	require.ErrorIs(t, card.Validate(), model.ErrCardHolderDocumentIsRequired, "Card holder document should be required")
}

func TestCardShouldUnmarshal(t *testing.T) {
	card := model.NewCard()
	err := card.Unmarshal([]byte(`{"number":"1234567890123456","security_code":"123","exp_month":"12","exp_year":"2025","holder":{"holder_name":"John Doe","holder_document":"12345678901"},"id_card":"1","id_merchant":"1","brand":"Visa","first_digits":"1234","last_digits":"5678","valid_until":"2025-12-31","created_at":"2023-01-01T00:00:00Z","updated_at":"2023-01-01T00:00:00Z"}`))
	require.NoError(t, err, "Card should unmarshal without error")
	require.Equal(t, "1234567890123456", card.Number, "Card number should match")
	require.Equal(t, "123", card.CVV, "Card security code should match")
	require.Equal(t, "12", card.ExpirationMonth, "Card expiration month should match")
	require.Equal(t, "2025", card.ExpirationYear, "Card expiration year should match")
	require.Equal(t, "John Doe", card.Holder.Name, "Card holder name should match")
	require.Equal(t, "12345678901", card.Holder.Document, "Card holder document should match")
	require.Equal(t, "1", card.ID, "Card ID should match")
	require.Equal(t, "1", card.MerchantID, "Card merchant ID should match")
	require.Equal(t, "Visa", card.Brand, "Card brand should match")
	require.Equal(t, "1234", card.FirstDigits, "Card first digits should match")
	require.Equal(t, "5678", card.LastDigits, "Card last digits should match")
	require.Equal(t, "2025-12-31", card.ValidUntil, "Card valid until should match")
	require.Equal(t, "2023-01-01T00:00:00Z", card.CreatedAt, "Card created at should match")
	require.Equal(t, "2023-01-01T00:00:00Z", card.UpdatedAt, "Card updated at should match")
	require.NoError(t, card.Validate(), "Card should be valid after unmarshalling")
}
