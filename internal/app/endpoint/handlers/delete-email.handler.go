package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
	"github.com/quadgod/email-service-go/internal/app/usecase"
	log "github.com/sirupsen/logrus"
)

func BuildDeleteEmailHandler(deleteEmailUseCase usecase.IDeleteEmailUseCase) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := deleteEmailUseCase.Delete(id)

		if err != nil {
			if err.Error() == repository.EmailNotFoundError {
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
