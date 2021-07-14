package repos

import (
	"context"
	"errors"
	"time"

	"github.com/quadgod/email-service-go/internal/app/db"
	"github.com/quadgod/email-service-go/internal/app/db/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const EmailNotFoundError = "EMAIL_NOT_FOUND_ERROR"

type IEmailRepository interface {
	Insert(email *entities.Email) (*entities.Email, error)
	Commit(id string) (*entities.Email, error)
	Delete(id string) error
	GetEmailForSend() (*entities.Email, error)
	MarkEmailAsSent(id string) (*entities.Email, error)
	UnlockEmails() (int64, error)
}

type MongoEmailRepository struct {
	client *db.IMongoClient
}

const MongoNoDocumentsInResultError = "mongo: no documents in result"

func NewMongoEmailRepository(client *db.IMongoClient) IEmailRepository {
	return &MongoEmailRepository{
		client,
	}
}

func getEmailsCollection(client *db.IMongoClient) (*mongo.Collection, error) {
	database, dbErr := (*client).GetDatabase()
	if dbErr != nil {
		return nil, dbErr
	}

	emailsCollection := database.Collection("emails")
	return emailsCollection, nil
}

func (instance *MongoEmailRepository) UnlockEmails() (int64, error) {
	emailsCollection, dbErr := getEmailsCollection(instance.client)
	if dbErr != nil {
		return 0, dbErr
	}

	updateResult, err := emailsCollection.UpdateMany(
		context.TODO(),
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

func (instance *MongoEmailRepository) GetEmailForSend() (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(instance.client)
	if dbErr != nil {
		return nil, dbErr
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	updateResult := emailsCollection.FindOneAndUpdate(context.TODO(),
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
		if updateResult.Err().Error() == MongoNoDocumentsInResultError {
			return nil, errors.New(EmailNotFoundError)
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

func (instance *MongoEmailRepository) Delete(id string) error {
	emailsCollection, dbErr := getEmailsCollection(instance.client)
	if dbErr != nil {
		return dbErr
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, deleteErr := emailsCollection.DeleteOne(context.TODO(),
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

func (instance *MongoEmailRepository) Commit(id string) (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(instance.client)
	if dbErr != nil {
		return nil, dbErr
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := emailsCollection.FindOneAndUpdate(context.TODO(),
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
		if updateResult.Err().Error() == MongoNoDocumentsInResultError {
			return nil, errors.New(EmailNotFoundError)
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

func (instance *MongoEmailRepository) Insert(newEmail *entities.Email) (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(instance.client)
	if dbErr != nil {
		return nil, dbErr
	}

	insertResult, insertErr := emailsCollection.InsertOne(context.TODO(), newEmail)

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

func (instance *MongoEmailRepository) MarkEmailAsSent(id string) (*entities.Email, error) {
	emailsCollection, dbErr := getEmailsCollection(instance.client)
	if dbErr != nil {
		return nil, dbErr
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := emailsCollection.FindOneAndUpdate(context.TODO(),
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
		if updateResult.Err().Error() == MongoNoDocumentsInResultError {
			return nil, errors.New(EmailNotFoundError)
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
