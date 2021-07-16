package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func ConnectMongoClient(parentContext context.Context, client *mongo.Client) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(parentContext, 10*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, err
}
