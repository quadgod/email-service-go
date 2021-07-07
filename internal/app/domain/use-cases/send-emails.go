package usecases

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	emailrepository "github.com/quadgod/email-service-go/internal/app/db/repositories/email.repository"
	"github.com/quadgod/email-service-go/internal/app/domain/providers"
	log "github.com/sirupsen/logrus"
)

type ISendEmailsUseCase interface {
	StartSending()
}

type SendEmailsUseCase struct {
	emailRepository emailrepository.IEmailRepository
	config          config.IConfig
}

// func sendEmail(emailRepository repos.IEmailRepository) {

// }

// func sendEmails(emailRepository repos.IEmailRepository, config config.IConfig) {
// 	var counter int64 = 0
// 	var shouldSleep = false
// 	var max = config.GetMaxEmailsPerInterval()

// 	for counter < max {

// 		counter++
// 	}
// }

func sendEmail(
	config config.IConfig,
	emailRepository emailrepository.IEmailRepository,
	emailProvider providers.IEmailProvider,
	sentChannel chan string,
) {
	emailForSend, emailRepoErr := emailRepository.GetEmailForSend()

	if emailRepoErr != nil {
		if emailRepoErr.Error() == emailrepository.ERROR_EMAIL_NOT_FOUND {
			log.Debug("No emails for send found")
		} else {
			log.Error("Get email for send error", emailRepoErr.Error())
			close(sentChannel)
		}
		return
	}

	providerErr := emailProvider.Send(emailForSend)
	if providerErr != nil {
		log.Errorf("Email provider error: %d %s", providerErr.Status, providerErr.Message)
		close(sentChannel)
		return
	}

	sentEmail, markErr := emailRepository.MarkEmailAsSent(emailForSend.ID.String())
}

// func (instance SendEmailsUseCase) StartSending() {

// 	for counter <

// 	email, err := instance.EmailRepository.GetEmalForSend()
// 	if err != nil && err.Error() == constants.ERROR_NOT_FOUND {
// 		counter = 0
// 		start = nil
// 		return
// 	}

// 	var emailsSent = 0
// 	var rateLimitInterval = instance.Config.GetRateLimitIntervalMs()
// 	var maxEmailsPerInterval = instance.Config.GetMaxEmailsPerInterval()

// 	sendEmail := func() {
// 		email, err := instance.EmailRepository.GetEmalForSend()
// 		if err != nil && err.Error() != constants.ERROR_NOT_FOUND {

// 		}
// 		fmt.Println("")
// 	}

// 	sendEmails()

// }
