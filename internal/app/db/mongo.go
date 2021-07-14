package db

import (
	"context"
	"github.com/quadgod/email-service-go/internal/app/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type IMongoClient interface {
	Connect() (*mongo.Client, error)
	GetDatabase() (*mongo.Database, error)
}

type MongoClient struct {
	sync.Once
	clientError *error
	client      *mongo.Client
	config      *config.IConfig
}

func NewMongoClient(
	config *config.IConfig,
) IMongoClient {
	return &MongoClient{client: nil, config: config}
}

func (instance *MongoClient) GetDatabase() (*mongo.Database, error) {
	client, err := instance.Connect()
	if err != nil {
		return nil, err
	}

	database := client.Database((*instance.config).GetDatabaseName())
	return database, nil
}

func (instance *MongoClient) Connect() (*mongo.Client, error) {
	instance.Do(func() {
		opts := options.
			Client().
			SetMaxPoolSize(50).
			SetHeartbeatInterval(5000).
			ApplyURI((*instance.config).GetDbUrl())

		client, err := mongo.NewClient(opts)
		if err != nil {
			log.Error("Create mongo client instance error", err)
			instance.clientError = &err
			return
		}

		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Error("Connect mongo", err)
			instance.clientError = &err
			return
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			instance.clientError = &err
			return
		}

		instance.client = client
	})

	if instance.clientError != nil {
		return nil, *instance.clientError
	}

	return instance.client, nil
}
