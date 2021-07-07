package providers

import "github.com/quadgod/email-service-go/internal/app/db/entities"

type EmailProviderError struct {
	Status  int
	Message string
}

type IEmailProvider interface {
	Send(email *entities.Email) *EmailProviderError
}
