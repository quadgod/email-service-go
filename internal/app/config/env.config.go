package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getEnvValueAsString(key string) string {
	val, ok := os.LookupEnv(key)

	if ok {
		return strings.TrimSpace(val)
	} else {
		return ""
	}
}

func getEnvValueAsStringOrDefault(key string, defaultValue string) string {
	value := getEnvValueAsString(key)

	if value == "" {
		return defaultValue
	}

	return value
}

type EnvConfig struct{}

func NewEnvConfig() IConfig {
	var config IConfig = &EnvConfig{}
	return config
}

func (config EnvConfig) GetRateLimitIntervalMs() int64 {
	interval, err := strconv.ParseInt(getEnvValueAsStringOrDefault("RATE_LIMIT_INTERVAL_MS", "60000"), 10, 64)
	if err != nil {
		return 60000
	}
	return interval
}

func (config EnvConfig) GetMaxEmailsPerInterval() int64 {
	interval, err := strconv.ParseInt(getEnvValueAsStringOrDefault("MAX_EMAILS_PER_INTERVAL", "10"), 10, 64)
	if err != nil {
		return 1000
	}
	return interval
}

func (config EnvConfig) GetAppPort() string {
	return getEnvValueAsStringOrDefault("APP_PORT", "3000")
}

func (config EnvConfig) GetDatabseName() string {
	return getEnvValueAsStringOrDefault("DB_NAME", "sandbox")
}

func (config EnvConfig) GetDatabaseUser() string {
	return getEnvValueAsStringOrDefault("DB_USER", "test")
}

func (config EnvConfig) GetDatabaseUserPassword() string {
	return getEnvValueAsStringOrDefault("DB_USER_PASSWORD", "test")
}

func (config EnvConfig) GetDatabaseProtocol() string {
	return getEnvValueAsStringOrDefault("DB_PROTOCOL", "mongodb+srv")
}

func (config EnvConfig) GetDatabaseHost() string {
	return getEnvValueAsStringOrDefault("DB_HOST", "cluster0.ipkhc.mongodb.net")
}

func (config EnvConfig) GetDbUrl() string {
	dbProtocol := config.GetDatabaseProtocol()
	dbUser := config.GetDatabaseUser()
	dbPassword := config.GetDatabaseUserPassword()
	dbHost := config.GetDatabaseHost()
	dbName := config.GetDatabseName()
	dbUrl := fmt.Sprintf("%s://%s:%s@%s/%s?retryWrites=true&w=majority", dbProtocol, dbUser, dbPassword, dbHost, dbName)
	return dbUrl
}
