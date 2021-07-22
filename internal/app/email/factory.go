package email

import (
	"errors"
	"github.com/quadgod/email-service-go/internal/app/config"
)

type ProviderFactory struct {
	config *config.Config
}

func NewFactory(config *config.Config) ProvidersFactory {
	return &ProviderFactory{config: config}
}

func (e *ProviderFactory) GetProviderByName(providerName string) (Provider, error) {
	switch providerName {
	case "fake":
		return NewFakeProvider(e.config), nil
	default:
		return nil, errors.New(UnknownProviderError)
	}
}
