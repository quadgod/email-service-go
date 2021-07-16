package usecase

import (
	"context"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
)

type IDeleteEmailUseCase interface {
	Delete(ctx context.Context, id string) error
}

type DeleteEmailUseCase struct {
	emailRepository *repository.IEmailRepository
}

func NewDeleteEmailUseCase(emailRepository *repository.IEmailRepository) IDeleteEmailUseCase {
	return &DeleteEmailUseCase{
		emailRepository,
	}
}

func (instance DeleteEmailUseCase) Delete(ctx context.Context, id string) error {
	err := (*instance.emailRepository).Delete(ctx, id)
	return err
}
