package usecases

import (
	"github.com/quadgod/email-service-go/internal/app/db/entities"
	"github.com/quadgod/email-service-go/internal/app/db/repos"
)

type IEmailCommitter interface {
	Commit(id string) (*entities.Email, error)
}

type EmailCommitter struct {
	EmailRepository repos.IEmailRepository
}

func (instance EmailCommitter) Commit(id string) (*entities.Email, error) {
	entity, err := instance.EmailRepository.Commit(id)
	return entity, err
}
