package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/quadgod/email-service-go/internal/app/domain/use-cases"
)

func BuildCreateEmailHandler(emailCreator usecases.IEmailCreator) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload usecases.CreateEmailDTO

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email, createErr := emailCreator.Create(payload)

		if createErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
			return
		}

		c.JSON(http.StatusOK, email)
	}
}
