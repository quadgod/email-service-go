package emailrepository

import (
	"context"
	"errors"
	"time"

	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/entities"
	"github.com/quadgod/email-service-go/internal/app/db/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoEmailRepository struct {
	config config.IConfig
}

const mongoNoDocumentsInResultError = "mongo: no documents in result"

func NewMongoEmailRepository(config config.IConfig) IEmailRepository {
	repo := &MongoEmailRepository{
		config: config,
	}

	return repo
}

func getEmailsCollection(config config.IConfig) (*mongo.Collection, error) {
	database, dbErr := utils.GetMongoDatabase(config)
	if dbErr != nil {
		return nil, dbErr
	}

	emailsCollection := database.Collection("emails")
	return emailsCollection, nil
}

func (repo MongoEmailRepository) GetEmailForSend() (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(repo.config)
	if dbErr != nil {
		return nil, dbErr
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	updateResult := emailsCollection.FindOneAndUpdate(context.Background(),
		bson.M{
			"$or": bson.A{
				bson.M{"lockedAt": bson.M{"$exists": false}},
				bson.M{"lockedAt": bson.M{"$eq": bsontype.Null}},
			},
			"readyToSend": true,
		},
		bson.M{
			"$set": bson.M{
				"lockedAt": time.Now(),
			},
		},
		&opt,
	)

	if updateResult.Err() != nil {
		if updateResult.Err().Error() == mongoNoDocumentsInResultError {
			return nil, errors.New(ERROR_EMAIL_NOT_FOUND)
		}

		return nil, updateResult.Err()
	}

	var email *entities.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}

func (repo MongoEmailRepository) Delete(id string) error {
	emailsCollection, dbErr := getEmailsCollection(repo.config)
	if dbErr != nil {
		return dbErr
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, deleteErr := emailsCollection.DeleteOne(context.Background(),
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
		return errors.New(ERROR_EMAIL_NOT_FOUND)
	}

	return nil
}

func (repo MongoEmailRepository) Commit(id string) (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(repo.config)
	if dbErr != nil {
		return nil, dbErr
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := emailsCollection.FindOneAndUpdate(context.Background(),
		bson.M{
			"_id": objId,
			"readyToSend": bson.M{
				"$ne": true,
			},
		},
		bson.M{
			"$set": bson.M{
				"readyToSend": true,
				"committedAt": time.Now(),
			},
		},
		&opt)

	if updateResult.Err() != nil {
		if updateResult.Err().Error() == mongoNoDocumentsInResultError {
			return nil, errors.New(ERROR_EMAIL_NOT_FOUND)
		}

		return nil, updateResult.Err()
	}

	var email *entities.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}

func (repo MongoEmailRepository) Insert(newEmail entities.Email) (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(repo.config)
	if dbErr != nil {
		return nil, dbErr
	}

	insertResult, insertErr := emailsCollection.InsertOne(context.Background(), newEmail)

	if insertErr != nil {
		return nil, insertErr
	}

	var email *entities.Email
	result := emailsCollection.FindOne(context.TODO(), bson.M{"_id": insertResult.InsertedID})

	if result.Err() != nil {
		return nil, result.Err()
	}

	decodeErr := result.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}

func (repo MongoEmailRepository) MarkEmailAsSent(id string) (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(repo.config)
	if dbErr != nil {
		return nil, dbErr
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := emailsCollection.FindOneAndUpdate(context.Background(),
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
			return nil, errors.New(ERROR_EMAIL_NOT_FOUND)
		}

		return nil, updateResult.Err()
	}

	var email *entities.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return email, nil
}
