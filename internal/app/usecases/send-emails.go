package usecases

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entities"
	"github.com/quadgod/email-service-go/internal/app/db/repos"
	"github.com/quadgod/email-service-go/internal/app/providers"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ISendEmailsUseCase interface {
	StartSending()
}

type SendEmailsUseCase struct {
	emailProvider   *providers.IEmailProvider
	emailRepository *repos.IEmailRepository
	config          *config.IConfig
}

func NewSendEmailsUseCase(
	emailProvider *providers.IEmailProvider,
	emailRepository *repos.IEmailRepository,
	config *config.IConfig,
) ISendEmailsUseCase {
	return &SendEmailsUseCase{
		emailProvider,
		emailRepository,
		config,
	}
}

func sendEmail(
	sending *sync.WaitGroup,
	email *entities.Email,
	emailRepository *repos.IEmailRepository,
	emailProvider *providers.IEmailProvider,
) {
	defer sending.Done()

	emailProviderErr := (*emailProvider).Send(email)
	if emailProviderErr != nil {
		if emailProviderErr.Error() == providers.MaxRequestRateLimitExceededError {
			log.Errorf("Email provider error: Maximum request rate limit has been exceeded for provider \"%s\"", email.Provider)
		} else {
			log.Errorf("Email provider error: %s", emailProviderErr.Error())
		}
		return
	}

	email, markError := (*emailRepository).MarkEmailAsSent(email.ID.String())
	if markError != nil {
		if markError.Error() == repos.EmailNotFoundError {
			log.Error("Mark email as sent error: Email not found")
		} else {
			log.Errorf("Mark email as sent error: %s", markError.Error())
		}
		return
	}

	log.Infof("Email id=\"%s\" sent at %s", email.ID.String(), email.SentAt.Local().String())
}

func sendEmails(
	allDone *sync.WaitGroup,
	emailRepository *repos.IEmailRepository,
	emailProvider *providers.IEmailProvider,
) {
	defer allDone.Done()

	var sending sync.WaitGroup

	for {
		email, err := (*emailRepository).GetEmailForSend()
		if err != nil {
			if err.Error() == repos.EmailNotFoundError {
				log.Info("No emails for send found")
			} else {
				log.Errorf("Get emails for send error: %s", err.Error())
			}
			break
		}
		log.Infof("Email for send found %s", email.ID.String())
		sending.Add(1)
		go sendEmail(&sending, email, emailRepository, emailProvider)
	}

	sending.Wait()
}

func (instance *SendEmailsUseCase) StartSending() {
	var allDone sync.WaitGroup
	for {
		log.Info("Starting new send emails iteration")
		allDone.Add(1)
		go sendEmails(&allDone, instance.emailRepository, instance.emailProvider)
		allDone.Wait()
		sleepInterval := (*instance.config).GetSendSleepIntervalSec()
		log.Info("Send emails iteration done")
		log.Infof("Going to sleep for %d seconds", sleepInterval)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
