package emailrepository

import "github.com/quadgod/email-service-go/internal/app/db/entities"

const ERROR_EMAIL_NOT_FOUND = "ERROR_EMAIL_NOT_FOUND"

type IEmailRepository interface {
	Insert(email entities.Email) (*entities.Email, error)
	Commit(id string) (*entities.Email, error)
	Delete(id string) error
	GetEmailForSend() (*entities.Email, error)
	MarkEmailAsSent(id string) (*entities.Email, error)
}
