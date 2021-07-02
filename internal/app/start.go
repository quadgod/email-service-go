package app

import (
	"fmt"
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/endpoints"
)

func Start() {
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
