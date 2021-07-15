package usecase

import (
	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
)

type ICommitEmailUseCase interface {
	Commit(id string) (*entity.Email, error)
}

type CommitEmailUseCase struct {
	emailRepository *repository.IEmailRepository
}

func NewCommitEmailUseCase(emailRepository *repository.IEmailRepository) ICommitEmailUseCase {
	return &CommitEmailUseCase{
		emailRepository,
	}
}

func (instance *CommitEmailUseCase) Commit(id string) (*entity.Email, error) {
	entity, err := (*instance.emailRepository).Commit(id)
	return entity, err
}
