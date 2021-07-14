package providers

import "github.com/quadgod/email-service-go/internal/app/db/entities"

const MaxRequestRateLimitExceededError = "MAX_REQUEST_RATE_LIMIT_EXCEEDED_ERROR"

type IEmailProvider interface {
	Send(email *entities.Email) error
}
