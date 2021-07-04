package app

import (
	"github.com/quadgod/email-service-go/internal/app/db/repos"
	usecases "github.com/quadgod/email-service-go/internal/app/domain/use-cases"
)

var EmailRepository repos.IEmailRepository = &repos.MongoEmailRepository{}
var EmailCreator usecases.IEmailCreator = &usecases.EmailCreator{
	EmailRepository: EmailRepository,
}
var EmailCommitter usecases.IEmailCommitter = &usecases.EmailCommitter{
	EmailRepository: EmailRepository,
}
