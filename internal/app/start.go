package app

import (
	"fmt"
	"github.com/quadgod/email-service-go/internal/app/providers"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db"
	"github.com/quadgod/email-service-go/internal/app/db/repos"
	"github.com/quadgod/email-service-go/internal/app/endpoints"
	"github.com/quadgod/email-service-go/internal/app/usecases"
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

	mongoClient := db.NewMongoClient(&envConfig)
	_, mongoConnectionError := mongoClient.Connect()
	if mongoConnectionError != nil {
		panic(mongoConnectionError)
	}

	emailRepository := repos.NewMongoEmailRepository(&mongoClient)
	createEmailUseCase := usecases.NewCreateEmailUseCase(&emailRepository)
	commitEmailUseCase := usecases.NewCommitEmailUseCase(&emailRepository)
	deleteEmailUseCase := usecases.NewDeleteEmailUseCase(&emailRepository)
	emailProvider := providers.NewFakeEmailProvider(&envConfig)
	sendEmailsUseCase := usecases.NewSendEmailsUseCase(
		&emailProvider,
		&emailRepository,
		&envConfig,
	)

	go sendEmailsUseCase.StartSending()

	router := gin.Default()

	endpoints.Setup(
		router,
		createEmailUseCase,
		commitEmailUseCase,
		deleteEmailUseCase,
	)

	port := envConfig.GetAppPort()
	err := endless.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email service stopped")
	os.Exit(0)
}
