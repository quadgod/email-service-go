package config

import (
	"fmt"
	"os"
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

func GetAppPort() string {
	return getEnvValueAsStringOrDefault("APP_PORT", "3000")
}

func GetDatabseName() string {
	return getEnvValueAsStringOrDefault("DB_NAME", "sandbox")
}

func GetDatabaseUser() string {
	return getEnvValueAsStringOrDefault("DB_USER", "test")
}

func GetDatabaseUserPassword() string {
	return getEnvValueAsStringOrDefault("DB_USER_PASSWORD", "test")
}

func GetDatabaseProtocol() string {
	return getEnvValueAsStringOrDefault("DB_PROTOCOL", "mongodb+srv")
}

func GetDatabaseHost() string {
	return getEnvValueAsStringOrDefault("DB_HOST", "cluster0.ipkhc.mongodb.net")
}

func GetDbUrl() string {
	dbProtocol := GetDatabaseProtocol()
	dbUser := GetDatabaseUser()
	dbPassword := GetDatabaseUserPassword()
	dbHost := GetDatabaseHost()
	dbName := GetDatabseName()
	dbUrl := fmt.Sprintf("%s://%s:%s@%s/%s?retryWrites=true&w=majority", dbProtocol, dbUser, dbPassword, dbHost, dbName)
	return dbUrl
}
