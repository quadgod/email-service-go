package usecase

import (
	"context"
	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
)

type ICommitEmailUseCase interface {
	Commit(ctx context.Context, id string) (*entity.Email, error)
}

type CommitEmailUseCase struct {
	emailRepository *repository.IEmailRepository
}

func NewCommitEmailUseCase(emailRepository *repository.IEmailRepository) ICommitEmailUseCase {
	return &CommitEmailUseCase{
		emailRepository,
	}
}

func (instance *CommitEmailUseCase) Commit(ctx context.Context, id string) (*entity.Email, error) {
	email, err := (*instance.emailRepository).Commit(ctx, id)
	return email, err
}
