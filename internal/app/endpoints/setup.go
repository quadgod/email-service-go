package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/endpoints/handlers"
)

func Setup(router *gin.Engine) {
	router.POST("/", handlers.CreateEmail)
}
