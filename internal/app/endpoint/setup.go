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
	router.POST("/", handlers.NewCreateEmailHandler(createEmailUseCase))
	router.PATCH("/:id/commit", handlers.NewCommitEmailHandler(commitEmailUseCase))
	router.DELETE("/:id", handlers.NewDeleteEmailHandler(deleteEmailUseCase))
}
