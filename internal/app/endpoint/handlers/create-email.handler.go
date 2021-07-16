package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/usecase"
	log "github.com/sirupsen/logrus"
)

func BuildCreateEmailHandler(createEmailUseCase usecase.ICreateEmailUseCase) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload usecase.CreateEmailDTO

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email, createErr := createEmailUseCase.Create(c, &payload)

		if createErr != nil {
			log.Error("[create email]: Internal server error", createErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, email)
	}
}
