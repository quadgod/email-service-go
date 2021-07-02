package db

import (
	"context"

	"github.com/quadgod/email-service-go/internal/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var ctx context.Context = context.Background()

func GetMongoClient() (*mongo.Client, error) {
	if mongoClient != nil {
		return mongoClient, nil
	}

	var err error
	mongoClient, err = mongo.NewClient(options.Client().SetHeartbeatInterval(5000).ApplyURI(config.GetDbUrl()))

	if err != nil {
		return nil, err
	}

	err = mongoClient.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return mongoClient, err
}
