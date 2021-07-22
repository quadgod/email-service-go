package email

import "github.com/quadgod/email-service-go/internal/app/db/entity"

type Provider interface {
	Send(email *entity.Email) error
}

type ProvidersFactory interface {
	GetProviderByName(providerName string) (Provider, error)
}
