package utils

import (
	"context"
	"sync"

	"github.com/quadgod/email-service-go/internal/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error

var buildClientOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
	buildClientOnce.Do(func() {
		client, err := mongo.NewClient(options.Client().SetHeartbeatInterval(5000).ApplyURI(config.GetDbUrl()))
		if err != nil {
			clientInstanceError = err
			return
		}

		err = client.Connect(context.Background())
		if err != nil {
			clientInstanceError = err
			return
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			clientInstanceError = err
			return
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
