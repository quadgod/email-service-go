package providers

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entities"
)

type FakeEmailProvider struct {
	config *config.IConfig
}

func NewFakeEmailProvider(config *config.IConfig) IEmailProvider {
	return &FakeEmailProvider{
		config,
	}
}

func (provider *FakeEmailProvider) Send(email *entities.Email) error {
	if email.Provider != "internal" {
		return errors.New(fmt.Sprintf("Client for \"%s\" provider not implemented", email.Provider))
	}

	sleepTime := rand.Intn(300)
	time.Sleep(time.Duration(sleepTime)) // Simulate some long operation

	if sleepTime > 280 {
		return errors.New(MaxRequestRateLimitExceededError)
	}

	return nil
}
