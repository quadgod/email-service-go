package repos

import (
	"context"
	"time"

	"github.com/quadgod/email-service-go/internal/app/db/entities"
	"github.com/quadgod/email-service-go/internal/app/db/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IEmailRepository interface {
	Insert(email entities.Email) (*entities.Email, error)
	Commit(id string) (*entities.Email, error)
}

type MongoEmailRepository struct{}

func (repo MongoEmailRepository) Commit(id string) (*entities.Email, error) {
	database, dbErr := utils.GetAppDatabase()
	if dbErr != nil {
		return nil, dbErr
	}

	emailsCollection := database.Collection("emails")

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after, // Return new document after update
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	updateResult := emailsCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{
		"readyToSend": true,
		"committedAt": time.Now(),
	}}, &opt)

	if updateResult.Err() != nil {
		log.Error("[MongoEmailRepository.CommitEmail]: FindOneAndUpdate() error", updateResult.Err())
		return nil, updateResult.Err()
	}

	var email *entities.Email
	decodeErr := updateResult.Decode(&email)

	if decodeErr != nil {
		log.Error("[MongoEmailRepository.CommitEmail]: updateResult.Decode(&email) error", decodeErr)
		return nil, decodeErr
	}

	return email, nil
}

func (repo MongoEmailRepository) Insert(newEmail entities.Email) (*entities.Email, error) {
	database, dbErr := utils.GetAppDatabase()
	if dbErr != nil {
		return nil, dbErr
	}

	emailsCollection := database.Collection("emails")
	insertResult, insertErr := emailsCollection.InsertOne(context.Background(), newEmail)

	if insertErr != nil {
		log.Error("[MongoEmailRepository.InsertEmail]: InsertOne() error", insertErr)
		return nil, insertErr
	}

	var email *entities.Email
	result := emailsCollection.FindOne(context.TODO(), bson.M{"_id": insertResult.InsertedID})

	if result.Err() != nil {
		log.Error("[MongoEmailRepository.InsertEmail]: FindOne() error", result.Err())
		return nil, result.Err()
	}

	decodeErr := result.Decode(&email)

	if decodeErr != nil {
		log.Error("[MongoEmailRepository.InsertEmail]: result.Decode(&email) error", decodeErr)
		return nil, decodeErr
	}

	return email, nil
}
