package usecase

import "github.com/quadgod/email-service-go/internal/app/db/repository"

type IDeleteEmailUseCase interface {
	Delete(id string) error
}

type DeleteEmailUseCase struct {
	emailRepository *repository.IEmailRepository
}

func NewDeleteEmailUseCase(emailRepository *repository.IEmailRepository) IDeleteEmailUseCase {
	return &DeleteEmailUseCase{
		emailRepository,
	}
}

func (instance DeleteEmailUseCase) Delete(id string) error {
	err := (*instance.emailRepository).Delete(id)
	return err
}
