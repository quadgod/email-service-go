package db

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(config *config.Config) (*mongo.Client, error) {
	opts := options.
		Client().
		SetMaxPoolSize(50).
		SetHeartbeatInterval(5000).
		ApplyURI((*config).GetDbUrl())

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}
