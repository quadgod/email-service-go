package usecase

import (
	"context"
	"time"

	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
)

type CreateEmailDTO struct {
	Provider string `json:"provider" binding:"required,oneof=fake emarsys"`
	To       string `json:"to" binding:"required,email"`
	Cc       string `json:"cc" binding:"email"`
	From	 string `json:"from"`
	Subject  string `json:"subject" binding:"required"`
	Body     string `json:"body" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=html text"`
}

type ICreateEmailUseCase interface {
	Create(ctx context.Context, payload *CreateEmailDTO) (*entity.Email, error)
}

type CreateEmailUseCase struct {
	emailRepository *repository.IEmailRepository
}

func NewCreateEmailUseCase(emailRepository *repository.IEmailRepository) ICreateEmailUseCase {
	return &CreateEmailUseCase{
		emailRepository,
	}
}

func (instance *CreateEmailUseCase) Create(ctx context.Context, payload *CreateEmailDTO) (*entity.Email, error) {
	var now = time.Now()

	newEmail := entity.Email{
		Provider:    payload.Provider,
		To:          payload.To,
		Cc:          payload.Cc,
		From:		 payload.From,
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

	email, err := (*instance.emailRepository).Insert(ctx, &newEmail)
	return email, err
}
