package utils

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAppDatabase() (*mongo.Database, error) {
	client, err := GetMongoClient()

	if err != nil {
		log.Error("[GetAppDatabase]: error", err)
		return nil, err
	}

	return client.Database(config.GetDatabseName()), nil
}
