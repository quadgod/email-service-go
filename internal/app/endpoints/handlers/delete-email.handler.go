package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	emailrepository "github.com/quadgod/email-service-go/internal/app/db/repositories/email.repository"
	usecases "github.com/quadgod/email-service-go/internal/app/domain/use-cases"
	log "github.com/sirupsen/logrus"
)

func BuildDeleteEmailHandler(deleteEmailUseCase usecases.IDeleteEmailUseCase) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := deleteEmailUseCase.Delete(id)

		if err != nil {
			if err.Error() == emailrepository.ERROR_EMAIL_NOT_FOUND {
				log.Warn("[delete email]: Email not found")
				c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
				return
			}

			log.Error("[delete email]: Internal server error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"deleted": 1,
		})
	}
}
