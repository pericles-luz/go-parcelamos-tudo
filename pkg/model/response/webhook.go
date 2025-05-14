package response

import (
	"encoding/json"
	"errors"
)

var (
	ErrIDReceiverConfigIsRequired = errors.New("id_receiver_config is required")
)

type Webhook struct {
	IDReceiverConfig string `json:"id_receiver_config"`
	EstablishmentID  string `json:"establishment_id"`
	ContextName      string `json:"context_name"`
	EventName        string `json:"event_name"`
	URL              string `json:"url"`
	PrivateKey       string `json:"private_key"`
	CreatedAt        string `json:"created_at"`
	RemovedAt        string `json:"removed_at"`
}

func NewWebhook() *Webhook {
	return &Webhook{}
}

func (w *Webhook) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, w)
	if err != nil {
		return err
	}
	return w.Validate()
}

func (w *Webhook) Validate() error {
	if w.IDReceiverConfig == "" {
		return ErrIDReceiverConfigIsRequired
	}
	return nil
}
