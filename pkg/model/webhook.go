package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrContextNameIsRequired = errors.New("context_name is required")
	ErrEventNameIsRequired   = errors.New("event_name is required")
	ErrURLIsRequired         = errors.New("url is required")
	ErrPrivateKeyIsRequired  = errors.New("private_key is required")
)

type Webhook struct {
	ContextName string `json:"context_name"`
	EventName   string `json:"event_name"`
	URL         string `json:"url"`
	PrivateKey  string `json:"private_key"`
}

func NewWebhook() *Webhook {
	return &Webhook{}
}

func (w *Webhook) SetContextName(contextName string) *Webhook {
	w.ContextName = contextName
	return w
}

func (w *Webhook) SetEventName(eventName string) *Webhook {
	w.EventName = eventName
	return w
}

func (w *Webhook) SetURL(url string) *Webhook {
	w.URL = url
	return w
}

func (w *Webhook) SetPrivateKey(privateKey string) *Webhook {
	w.PrivateKey = privateKey
	return w
}

func (w *Webhook) Validate() error {
	if w.ContextName == "" {
		return ErrContextNameIsRequired
	}
	if w.EventName == "" {
		return ErrEventNameIsRequired
	}
	if w.URL == "" {
		return ErrURLIsRequired
	}
	if w.PrivateKey == "" {
		return ErrPrivateKeyIsRequired
	}
	return nil
}

func (w *Webhook) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, w); err != nil {
		return err
	}
	return nil
}

func (w *Webhook) ToAPI() map[string]interface{} {
	return map[string]interface{}{
		"context_name": w.ContextName,
		"event_name":   w.EventName,
		"url":          w.URL,
		"private_key":  w.PrivateKey,
	}
}
