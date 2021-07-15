package email

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entity"
)

type FakeProvider struct {
	config *config.IConfig
}

func NewFakeProvider(config *config.IConfig) IEmailProvider {
	return &FakeProvider{
		config,
	}
}

func (provider *FakeProvider) Send(email *entity.Email) error {
	sleepTime := rand.Intn(300)
	time.Sleep(time.Duration(sleepTime)) // Simulate some long operation
	fmt.Printf("Email %s sent", email.ID.String())
	if sleepTime > 280 {
		return errors.New(MaxRequestRateLimitExceededError)
	}

	return nil
}
