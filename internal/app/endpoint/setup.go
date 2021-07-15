package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/quadgod/email-service-go/internal/app/endpoint/handlers"
	"github.com/quadgod/email-service-go/internal/app/usecase"
)

func Setup(
	router *gin.Engine,
	createEmailUseCase usecase.ICreateEmailUseCase,
	commitEmailUseCase usecase.ICommitEmailUseCase,
	deleteEmailUseCase usecase.IDeleteEmailUseCase,
) {
	router.POST("/", handlers.BuildCreateEmailHandler(createEmailUseCase))
	router.PATCH("/:id/commit", handlers.BuildCommitEmailHandler(commitEmailUseCase))
	router.DELETE("/:id", handlers.BuildDeleteEmailHandler(deleteEmailUseCase))
}
