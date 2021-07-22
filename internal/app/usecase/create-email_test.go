package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"

	mocks "github.com/quadgod/email-service-go/internal/app/db/repository/mocks"
)

func TestCreateEmailUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	payload := &CreateEmailDTO{
		Provider: "fake",
		To: "test@to.com",
		Cc: "test@cc.com",
		From: "test@from.com",
		Subject: "subject",
		Body: "body",
		Type: "html",
	}

	emailToInsert := &entity.Email{
		Provider: payload.Provider,
		To: payload.To,
		Cc: payload.Cc,
		From: payload.From,
		Subject: payload.Subject,
		Body: payload.Body,
		Type: payload.Type,
		CreatedAt:   &now,
		LockedAt:    nil,
		SentAt:      nil,
		CommittedAt: nil,
		Attachments: make([]string, 0),
		ReadyToSend: false,
	}

	emailToReturn := *emailToInsert
	emailToReturn.ID = primitive.NewObjectID()

	repoMock := mocks.NewMockEmailRepository(ctrl)
	repoMock.EXPECT().GetTimeNow().Return(now)
	repoMock.EXPECT().Insert(nil, emailToInsert).Return(&emailToReturn, nil)
	subject := NewCreateEmailUseCase(repoMock)
	result, err := subject.Create(nil, payload)
	assert.Equal(t, err, nil)
	assert.Equal(t, &emailToReturn, result)
}
