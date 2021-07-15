package usecase

import (
	"time"

	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
)

type CreateEmailDTO struct {
	To       string `json:"to" binding:"required,email"`
	Provider string `json:"email" binding:"required,oneof=internal emarsys"`
	Cc       string `json:"cc" binding:"email"`
	Subject  string `json:"subject" binding:"required"`
	Body     string `json:"body" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=html text"`
}

type ICreateEmailUseCase interface {
	Create(payload *CreateEmailDTO) (*entity.Email, error)
}

type CreateEmailUseCase struct {
	emailRepository *repository.IEmailRepository
}

func NewCreateEmailUseCase(emailRepository *repository.IEmailRepository) ICreateEmailUseCase {
	return &CreateEmailUseCase{
		emailRepository,
	}
}

func (instance *CreateEmailUseCase) Create(payload *CreateEmailDTO) (*entity.Email, error) {
	var now = time.Now()

	newEmail := entity.Email{
		Provider:    payload.Provider,
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

	entity, err := (*instance.emailRepository).Insert(&newEmail)
	return entity, err
}
