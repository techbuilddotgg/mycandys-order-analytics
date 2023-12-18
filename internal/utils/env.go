package utils

import (
	"os"
)

func GetEnvVar(key string) (string, error) {

	value := os.Getenv(key)
	if len(value) == 0 {
		return "", nil
	}

	return value, nil
}
