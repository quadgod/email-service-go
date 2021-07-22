package usecase

import (
	"context"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
	eml "github.com/quadgod/email-service-go/internal/app/email"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ISendEmailsUseCase interface {
	Start()
}

type SendEmailsUseCase struct {
	providerFactory *eml.ProvidersFactory
	emailRepository *repository.EmailRepository
	config          *config.Config
}

func NewSendEmailsUseCase(
	providerFactory eml.ProvidersFactory,
	emailRepository repository.EmailRepository,
	config config.Config,
) ISendEmailsUseCase {
	return &SendEmailsUseCase{
		&providerFactory,
		&emailRepository,
		&config,
	}
}

func (s *SendEmailsUseCase) sendEmail(sending *sync.WaitGroup, email *entity.Email) {
	defer sending.Done()
	provider, err := (*s.providerFactory).GetProviderByName(email.Provider)
	if err != nil {
		log.Errorf("Email provider factory error: %s \"%s\"", err.Error(), email.Provider)
		return
	}

	emailProviderErr := provider.Send(email)
	if emailProviderErr != nil {
		if emailProviderErr.Error() == eml.MaxRequestRateLimitExceededError {
			log.Errorf("Email email error: Maximum request rate limit has been exceeded for email \"%s\"", email.Provider)
		} else {
			log.Errorf("Email email error: %s", emailProviderErr.Error())
		}
		return
	}

	email, markError := (*s.emailRepository).MarkAsSent(context.TODO(), email.ID.String())
	if markError != nil {
		if markError.Error() == repository.EmailNotFoundError {
			log.Error("Mark email as sent error: Email not found")
		} else {
			log.Errorf("Mark email as sent error: %s", markError.Error())
		}
		return
	}

	log.Infof("Email id=\"%s\" sent at %s", email.ID.String(), email.SentAt.Local().String())
}

func (s *SendEmailsUseCase) sendEmails(allDone *sync.WaitGroup) {
	defer allDone.Done()

	var sending sync.WaitGroup

	for {
		email, err := (*s.emailRepository).GetForSend(context.TODO())
		if err != nil {
			if err.Error() == repository.EmailNotFoundError {
				log.Info("No emails for send found")
			} else {
				log.Errorf("Get emails for send error: %s", err.Error())
			}
			break
		}
		log.Infof("Email for send found %s", email.ID.String())
		sending.Add(1)
		go s.sendEmail(&sending, email)
	}

	sending.Wait()
}

func (s *SendEmailsUseCase) Start() {
	var allDone sync.WaitGroup
	for {
		log.Info("Starting new send emails iteration")
		allDone.Add(1)
		go s.sendEmails(&allDone)
		allDone.Wait()
		sleepInterval := (*s.config).GetSendSleepIntervalSec()
		log.Info("Send emails iteration done")
		log.Infof("Going to sleep for %d seconds", sleepInterval)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
