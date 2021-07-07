package usecases

import (
	emailrepository "github.com/quadgod/email-service-go/internal/app/db/repositories/email.repository"
)

type IDeleteEmailUseCase interface {
	Delete(id string) error
}

type DeleteEmailUseCase struct {
	emailRepository emailrepository.IEmailRepository
}

func NewDeleteEmailUseCase(emailRepository emailrepository.IEmailRepository) IDeleteEmailUseCase {
	return &DeleteEmailUseCase{
		emailRepository,
	}
}

func (instance DeleteEmailUseCase) Delete(id string) error {
	err := instance.emailRepository.Delete(id)
	return err
}
