package utils

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoDatabase(config config.IConfig) (*mongo.Database, error) {
	client, err := GetMongoClient(config)

	if err != nil {
		return nil, err
	}

	return client.Database(config.GetDatabseName()), nil
}
