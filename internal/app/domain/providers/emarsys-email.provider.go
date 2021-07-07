package providers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entities"
)

type EmarsysEmailProvider struct {
	config config.IConfig
}

func NewEmarsysEmailProvider(config config.IConfig) IEmailProvider {
	return &EmarsysEmailProvider{
		config,
	}
}

// TODO: Impelement emarsys client
// https://dev.emarsys.com/v2/emarsys-developer-hub/whats-new
func (instance EmarsysEmailProvider) Send(email *entities.Email) *EmailProviderError {
	min := 300
	max := 100
	sleepTime := rand.Intn(max-min) + min
	time.Sleep(time.Duration(sleepTime)) // Simulate some long operation

	if sleepTime > 280 {
		return &EmailProviderError{
			Status:  http.StatusTooManyRequests,
			Message: "The maximum request rate limit has been exceeded 1000.",
		}
	}

	return nil
}
