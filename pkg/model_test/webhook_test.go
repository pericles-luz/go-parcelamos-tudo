package model_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestWebhookShouldValidate(t *testing.T) {
	webhook := model.NewWebhook()
	webhook.SetContextName("context_name").
		SetEventName("event_name").
		SetURL("https://example.com/webhook").
		SetPrivateKey("private_key")
	require.NoError(t, webhook.Validate())
}

func TestWebhookShouldNotValidateIfContextNameIsEmpty(t *testing.T) {
	webhook := model.NewWebhook()
	webhook.SetEventName("event_name").
		SetURL("https://example.com/webhook").
		SetPrivateKey("private_key")
	require.ErrorIs(t, webhook.Validate(), model.ErrContextNameIsRequired)
}

func TestWebhookShouldNotValidateIfEventNameIsEmpty(t *testing.T) {
	webhook := model.NewWebhook()
	webhook.SetContextName("context_name").
		SetURL("https://example.com/webhook").
		SetPrivateKey("private_key")
	require.ErrorIs(t, webhook.Validate(), model.ErrEventNameIsRequired)
}

func TestWebhookShouldNotValidateIfURLIsEmpty(t *testing.T) {
	webhook := model.NewWebhook()
	webhook.SetContextName("context_name").
		SetEventName("event_name").
		SetPrivateKey("private_key")
	require.ErrorIs(t, webhook.Validate(), model.ErrURLIsRequired)
}

func TestWebhookShouldNotValidateIfPrivateKeyIsEmpty(t *testing.T) {
	webhook := model.NewWebhook()
	webhook.SetContextName("context_name").
		SetEventName("event_name").
		SetURL("https://example.com/webhook")
	require.ErrorIs(t, webhook.Validate(), model.ErrPrivateKeyIsRequired)
}

func TestWebhookMustGenerateToMap(t *testing.T) {
	webhook := model.NewWebhook()
	webhook.SetContextName("context_name").
		SetEventName("event_name").
		SetURL("https://example.com/webhook").
		SetPrivateKey("private_key")
	require.Equal(t, map[string]interface{}{
		"context_name": "context_name",
		"event_name":   "event_name",
		"url":          "https://example.com/webhook",
		"private_key":  "private_key",
	}, webhook.ToAPI())
}

func TestWebhookMustUnmarshal(t *testing.T) {
	webhook := model.NewWebhook()
	require.NoError(t, webhook.Unmarshal([]byte(`{"context_name":"context_name","event_name":"event_name","url":"https://example.com/webhook","private_key":"private_key"}`)))
	require.Equal(t, "context_name", webhook.ContextName)
	require.Equal(t, "event_name", webhook.EventName)
	require.Equal(t, "https://example.com/webhook", webhook.URL)
	require.Equal(t, "private_key", webhook.PrivateKey)
	require.NoError(t, webhook.Validate())
	require.Equal(t, map[string]interface{}{
		"context_name": "context_name",
		"event_name":   "event_name",
		"url":          "https://example.com/webhook",
		"private_key":  "private_key",
	}, webhook.ToAPI())
}
