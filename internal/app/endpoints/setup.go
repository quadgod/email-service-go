package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/endpoints/handlers"
	"github.com/quadgod/email-service-go/internal/app/usecases"
)

func Setup(
	router *gin.Engine,
	createEmailUseCase usecases.ICreateEmailUseCase,
	commitEmailUseCase usecases.ICommitEmailUseCase,
	deleteEmailUseCase usecases.IDeleteEmailUseCase,
) {
	router.POST("/", handlers.BuildCreateEmailHandler(createEmailUseCase))
	router.PATCH("/:id/commit", handlers.BuildCommitEmailHandler(commitEmailUseCase))
	router.DELETE("/:id", handlers.BuildDeleteEmailHandler(deleteEmailUseCase))
}
