package config

import "github.com/quadgod/email-service-go/internal/app/utils"

func GetAppPort() string {
	return utils.GetEnvValueAsStringOrDefault("APP_PORT", "3000")
}

func GetDbUrl() string {
	return utils.GetEnvValueAsStringOrDefault("DB_URL", "mongodb://emails:emailspw@127.0.0.1/emails")
}
