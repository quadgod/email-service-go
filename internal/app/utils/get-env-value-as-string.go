package utils

import (
	"os"
	"strings"
)

func GetEnvValueAsString(key string) string {
	val, ok := os.LookupEnv(key)

	if ok {
		return strings.TrimSpace(val)
	} else {
		return ""
	}
}
