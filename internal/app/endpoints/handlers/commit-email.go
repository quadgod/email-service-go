package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/quadgod/email-service-go/internal/app/domain/use-cases"
)

func BuildCommitEmailHandler(emailCommitter usecases.IEmailCommitter) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		email, err := emailCommitter.Commit(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, email)
	}
}
