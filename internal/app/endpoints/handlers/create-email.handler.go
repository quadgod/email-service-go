package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/usecases"
	log "github.com/sirupsen/logrus"
)

func BuildCreateEmailHandler(createEmailUseCase usecases.ICreateEmailUseCase) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload usecases.CreateEmailDTO

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email, createErr := createEmailUseCase.Create(&payload)

		if createErr != nil {
			log.Error("[create email]: Internal server error", createErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, email)
	}
}
