package db

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEmailsCollection(client *mongo.Client, config *config.IConfig) *mongo.Collection {
	database := client.Database((*config).GetDatabaseName())
	collection := database.Collection("emails")
	return collection
}
