package usecases

import (
	"github.com/quadgod/email-service-go/internal/app/db/entities"
	"github.com/quadgod/email-service-go/internal/app/db/repos"
)

type ICommitEmailUseCase interface {
	Commit(id string) (*entities.Email, error)
}

type CommitEmailUseCase struct {
	emailRepository *repos.IEmailRepository
}

func NewCommitEmailUseCase(emailRepository *repos.IEmailRepository) ICommitEmailUseCase {
	return &CommitEmailUseCase{
		emailRepository,
	}
}

func (instance *CommitEmailUseCase) Commit(id string) (*entities.Email, error) {
	entity, err := (*instance.emailRepository).Commit(id)
	return entity, err
}
