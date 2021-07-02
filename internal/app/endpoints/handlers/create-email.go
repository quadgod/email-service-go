package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEmailDto struct {
	To      string `json:"to" binding:"required"`
	Cc      string `json:"cc"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func CreateEmail(c *gin.Context) {
	var body CreateEmailDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(200, "test")
}
