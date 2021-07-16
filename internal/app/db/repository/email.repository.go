package repository

import (
	"context"
	"errors"
	"time"

	"github.com/quadgod/email-service-go/internal/app/db/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const EmailNotFoundError = "EMAIL_NOT_FOUND_ERROR"
const mongoNoDocumentsInResultError = "mongo: no documents in result"

type IEmailRepository interface {
	Insert(ctx context.Context, email *entity.Email) (*entity.Email, error)
	Commit(ctx context.Context, id string) (*entity.Email, error)
	Delete(ctx context.Context, id string) error
	GetForSend(ctx context.Context) (*entity.Email, error)
	MarkAsSent(ctx context.Context, id string) (*entity.Email, error)
	Unlock(ctx context.Context) (int64, error)
}

type MongoEmailRepository struct {
	emailsCollection *mongo.Collection
}

func NewMongoEmailRepository(emailsCollection *mongo.Collection) IEmailRepository {
	return &MongoEmailRepository{
		emailsCollection,
	}
}

func (m *MongoEmailRepository) Unlock(ctx context.Context) (int64, error) {
	updateResult, err := m.emailsCollection.UpdateMany(
		ctx,
		bson.M{
			"$and": bson.A{
				bson.M{"lockedAt": bson.M{"$exists": true}},
				bson.M{"lockedAt": bson.M{"$ne": nil}},
				bson.M{
					"$or": bson.A{
						bson.M{"sentAt": bson.M{"$exists": false}},
						bson.M{"sentAt": bson.M{"$eq": nil}},
					},
				},
				bson.D{
					{"readyToSend", true},
					{"lockedAt", bson.M{
						"$lt": time.Now().Local().Add(-5 * time.Minute)},
					},
				},
			},
		},
		bson.M{
			"$set": bson.M{
				"lockedAt": nil,
			},
		},
	)

	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}

func (m *MongoEmailRepository) GetForSend(ctx context.Context) (*entity.Email, error) {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	updateResult := m.emailsCollection.FindOneAndUpdate(
		ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"lockedAt": bson.M{"$exists": false}},
				bson.M{"lockedAt": bson.M{"$eq": nil}},
			},
			"readyToSend": true,
		},
		bson.M{
			"$set": bson.M{
				"lockedAt": time.Now().Local(),
			},
		},
		&opt,
	)

	if updateResult.Err() != nil {
		if updateResult.Err().Error() == mongoNoDocumentsInResultError {
			return nil, errors.New(EmailNotFoundError)
		}

		return nil, updateResult.Err()
	}

	var email *entity.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}

func (m *MongoEmailRepository) Delete(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, deleteErr := m.emailsCollection.DeleteOne(
		ctx,
		bson.M{
			"_id": objId,
			"readyToSend": bson.M{
				"$ne": true,
			},
		},
	)

	if deleteErr != nil {
		return deleteErr
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New(EmailNotFoundError)
	}

	return nil
}

func (m *MongoEmailRepository) Commit(ctx context.Context, id string) (*entity.Email, error) {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := m.emailsCollection.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id": objId,
			"readyToSend": bson.M{
				"$ne": true,
			},
		},
		bson.M{
			"$set": bson.M{
				"readyToSend": true,
				"committedAt": time.Now().Local(),
			},
		},
		&opt)

	if updateResult.Err() != nil {
		if updateResult.Err().Error() == mongoNoDocumentsInResultError {
			return nil, errors.New(EmailNotFoundError)
		}

		return nil, updateResult.Err()
	}

	var email *entity.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}

func (m *MongoEmailRepository) Insert(ctx context.Context, newEmail *entity.Email) (*entity.Email, error) {
	insertResult, insertErr := m.emailsCollection.InsertOne(ctx, newEmail)

	if insertErr != nil {
		return nil, insertErr
	}

	var email *entity.Email
	result := m.emailsCollection.FindOne(ctx, bson.M{"_id": insertResult.InsertedID})

	if result.Err() != nil {
		return nil, result.Err()
	}

	decodeErr := result.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}

func (m *MongoEmailRepository) MarkAsSent(ctx context.Context, id string) (*entity.Email, error) {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := m.emailsCollection.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id": objId,
		},
		bson.M{
			"$set": bson.M{
				"sentAt": time.Now(),
			},
		},
		&opt)

	if updateResult.Err() != nil {
		if updateResult.Err().Error() == mongoNoDocumentsInResultError {
			return nil, errors.New(EmailNotFoundError)
		}

		return nil, updateResult.Err()
	}

	var email *entity.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}
