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

func NewEnvConfig() Config {
	return &EnvConfig{}
}

func (config *EnvConfig) GetLogLevel() string {
	return getEnvValueAsStringOrDefault("LOG_LEVEL", "debug")
}

func (config *EnvConfig) GetSendSleepIntervalSec() int {
	interval, err := strconv.ParseInt(getEnvValueAsStringOrDefault("SEND_SLEEP_INTERVAL_SECONDS", "60"), 10, 32)
	if err != nil {
		return 60
	}
	return int(interval)
}

func (config *EnvConfig) GetUnlockEmailsAfterSec() int {
	interval, err := strconv.ParseInt(getEnvValueAsStringOrDefault("UNLOCK_EMAILS_AFTER_SECONDS", "300"), 10, 32)
	if err != nil {
		return 300 // 5 mins
	}
	return int(interval)
}

func (config *EnvConfig) GetAppPort() string {
	return getEnvValueAsStringOrDefault("APP_PORT", "3000")
}

func (config *EnvConfig) GetDatabaseName() string {
	return getEnvValueAsStringOrDefault("DB_NAME", "sandbox")
}

func (config *EnvConfig) GetDatabaseUser() string {
	return getEnvValueAsStringOrDefault("DB_USER", "test")
}

func (config *EnvConfig) GetDatabaseUserPassword() string {
	return getEnvValueAsStringOrDefault("DB_USER_PASSWORD", "test")
}

func (config *EnvConfig) GetDatabaseProtocol() string {
	return getEnvValueAsStringOrDefault("DB_PROTOCOL", "mongodb+srv")
}

func (config *EnvConfig) GetDatabaseHost() string {
	return getEnvValueAsStringOrDefault("DB_HOST", "cluster0.ipkhc.mongodb.net")
}

func (config *EnvConfig) GetDbUrl() string {
	dbProtocol := config.GetDatabaseProtocol()
	dbUser := config.GetDatabaseUser()
	dbPassword := config.GetDatabaseUserPassword()
	dbHost := config.GetDatabaseHost()
	dbName := config.GetDatabaseName()

	if dbUser != "" && dbPassword != "" {
		return fmt.Sprintf("%s://%s:%s@%s/%s?retryWrites=true&w=majority", dbProtocol, dbUser, dbPassword, dbHost, dbName)
	}

	return fmt.Sprintf("%s://%s/%s?retryWrites=true&w=majority", dbProtocol, dbHost, dbName)
}
