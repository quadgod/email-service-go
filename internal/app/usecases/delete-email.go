package usecases

import "github.com/quadgod/email-service-go/internal/app/db/repos"

type IDeleteEmailUseCase interface {
	Delete(id string) error
}

type DeleteEmailUseCase struct {
	emailRepository *repos.IEmailRepository
}

func NewDeleteEmailUseCase(emailRepository *repos.IEmailRepository) IDeleteEmailUseCase {
	return &DeleteEmailUseCase{
		emailRepository,
	}
}

func (instance DeleteEmailUseCase) Delete(id string) error {
	err := (*instance.emailRepository).Delete(id)
	return err
}
