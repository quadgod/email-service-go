package repository

import (
	"context"
	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"time"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mocks.go

type EmailRepository interface {
	GetTimeNow() time.Time
	Insert(ctx context.Context, email *entity.Email) (*entity.Email, error)
	Commit(ctx context.Context, id string) (*entity.Email, error)
	Delete(ctx context.Context, id string) error
	GetForSend(ctx context.Context) (*entity.Email, error)
	MarkAsSent(ctx context.Context, id string) (*entity.Email, error)
	Unlock(ctx context.Context) (int64, error)
}