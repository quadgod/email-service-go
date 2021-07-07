package usecases

import (
	"github.com/quadgod/email-service-go/internal/app/db/entities"
	emailrepository "github.com/quadgod/email-service-go/internal/app/db/repositories/email.repository"
)

type ICommitEmailUseCase interface {
	Commit(id string) (*entities.Email, error)
}

type CommitEmailUseCase struct {
	emailRepository emailrepository.IEmailRepository
}

func NewCommitEmailUseCase(emailRepository emailrepository.IEmailRepository) ICommitEmailUseCase {
	commitEmailUseCase := &CommitEmailUseCase{
		emailRepository,
	}
	return commitEmailUseCase
}

func (instance CommitEmailUseCase) Commit(id string) (*entities.Email, error) {
	entity, err := instance.emailRepository.Commit(id)
	return entity, err
}
