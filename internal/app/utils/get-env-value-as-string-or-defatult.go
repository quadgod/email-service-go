package utils

func GetEnvValueAsStringOrDefault(key string, defaultValue string) string {
	value := GetEnvValueAsString(key)

	if value == "" {
		return defaultValue
	}

	return value
}
