package server

import (
	"fmt"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/endpoints"
	log "github.com/sirupsen/logrus"
)

func Start() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	router := gin.New()
	router.Use(gin.Recovery()) // Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.

	endpoints.Setup(router)

	port := config.GetAppPort()
	err := endless.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("email service stopped")
	os.Exit(0)
}
