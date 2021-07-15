package email

import (
	"errors"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entity"
)

type IEmailProvider interface {
	Send(email *entity.Email) error
}

type IProviderFactory interface {
	GetProviderByName(providerName string) (IEmailProvider, error)
}

type ProviderFactory struct {
	config *config.IConfig
}

func NewFactory(config *config.IConfig) IProviderFactory {
	return &ProviderFactory{config: config}
}

func (e *ProviderFactory) GetProviderByName(providerName string) (IEmailProvider, error) {
	switch providerName {
	case "fake":
		return NewFakeProvider(e.config), nil
	default:
		return nil, errors.New(UnknownProviderError)
	}
}
