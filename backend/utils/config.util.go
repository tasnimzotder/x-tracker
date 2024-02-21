package utils

import (
	"log"
	"os"
)

func GetEnvVariable(key string) string {
	key = "XT_" + key
	value, err := os.LookupEnv(key)

	if !err {
		log.Fatalf("Environment variable %s not set", key)
		return ""
	} else {
		return value
	}
}
