package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/db/repos"
	"github.com/quadgod/email-service-go/internal/app/usecases"
	log "github.com/sirupsen/logrus"
)

func BuildCommitEmailHandler(commitEmailUseCase usecases.ICommitEmailUseCase) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		email, err := commitEmailUseCase.Commit(id)

		if err != nil {
			if err.Error() == repos.EmailNotFoundError {
				log.Warn("[commit email]: Email not found")
				c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
				return
			}

			log.Error("[commit email]: Internal server error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, email)
	}
}
