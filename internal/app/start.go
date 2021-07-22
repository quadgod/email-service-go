package app

import (
	"context"
	"fmt"
	"github.com/quadgod/email-service-go/internal/app/email"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
	"github.com/quadgod/email-service-go/internal/app/endpoint"
	"github.com/quadgod/email-service-go/internal/app/usecase"
	log "github.com/sirupsen/logrus"
)

func Start() {
	envConfig := config.NewEnvConfig()

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	logLevel, logLevelParseError := log.ParseLevel(envConfig.GetLogLevel())

	if logLevelParseError != nil {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(logLevel)
	}

	client, err := db.CreateMongoClient(&envConfig)
	if err != nil {
		log.Fatal(err)
	}

	connectedClient, err := db.ConnectMongoClient(context.Background(), client)
	if err != nil {
		log.Fatal(err)
	}

	emailsCollection := db.GetEmailsCollection(connectedClient, &envConfig)

	emailRepository := repository.NewMongoEmailRepository(emailsCollection)
	createEmailUseCase := usecase.NewCreateEmailUseCase(emailRepository)
	commitEmailUseCase := usecase.NewCommitEmailUseCase(emailRepository)
	deleteEmailUseCase := usecase.NewDeleteEmailUseCase(emailRepository)
	emailProviderFactory := email.NewFactory(&envConfig)
	sendEmailsUseCase := usecase.NewSendEmailsUseCase(
		emailProviderFactory,
		emailRepository,
		envConfig,
	)
	unlockEmailsUseCase := usecase.NewUnlockEmailsUseCase(
		emailRepository,
		envConfig,
	)

	go sendEmailsUseCase.Start()
	go unlockEmailsUseCase.Start()

	router := gin.Default()

	endpoint.Setup(
		router,
		createEmailUseCase,
		commitEmailUseCase,
		deleteEmailUseCase,
	)

	port := envConfig.GetAppPort()
	listenError := endless.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if listenError != nil {
		log.Fatal(listenError)
	}

	log.Println("Email service stopped")
	os.Exit(0)
}
