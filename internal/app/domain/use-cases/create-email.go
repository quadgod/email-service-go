package usecases

import (
	"time"

	"github.com/quadgod/email-service-go/internal/app/db/entities"

	emailrepository "github.com/quadgod/email-service-go/internal/app/db/repositories/email.repository"
)

type CreateEmailDTO struct {
	To      string `json:"to" binding:"required,email"`
	Cc      string `json:"cc" binding:"email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
	Type    string `json:"type" binding:"required,oneof=html text"`
}

type ICreateEmailUseCase interface {
	Create(payload CreateEmailDTO) (*entities.Email, error)
}

type CreateEmailUseCase struct {
	emailRepository emailrepository.IEmailRepository
}

func NewCreateEmailUseCase(emailRepository emailrepository.IEmailRepository) ICreateEmailUseCase {
	return &CreateEmailUseCase{
		emailRepository,
	}
}

func (instance CreateEmailUseCase) Create(payload CreateEmailDTO) (*entities.Email, error) {
	var now = time.Now()

	newEmail := entities.Email{
		To:          payload.To,
		Cc:          payload.Cc,
		Subject:     payload.Subject,
		Body:        payload.Body,
		Type:        payload.Type,
		CreatedAt:   &now,
		LockedAt:    nil,
		SentAt:      nil,
		CommittedAt: nil,
		Attachments: make([]string, 0),
		ReadyToSend: false,
	}

	entity, err := instance.emailRepository.Insert(newEmail)
	return entity, err
}
