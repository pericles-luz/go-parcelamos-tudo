package response_test

import (
	"testing"

	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
	"github.com/stretchr/testify/require"
)

func TestWebhookShouldValidate(t *testing.T) {
	webhook := response.NewWebhook()
	webhook.IDReceiverConfig = "receiver_config_id"
	webhook.ContextName = "context_name"
	webhook.EventName = "event_name"
	webhook.URL = "https://example.com/webhook"
	webhook.PrivateKey = "private_key"
	require.NoError(t, webhook.Validate())
}

func TestWebhookShouldNotValidateIfIDReceiverConfigIsEmpty(t *testing.T) {
	webhook := response.NewWebhook()
	webhook.ContextName = "context_name"
	webhook.EventName = "event_name"
	webhook.URL = "https://example.com/webhook"
	webhook.PrivateKey = "private_key"
	require.ErrorIs(t, webhook.Validate(), response.ErrIDReceiverConfigIsRequired)
}
