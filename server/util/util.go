package util

import (
	"log"
	"os"
)

func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing mandatory environment variable: %s", key)
	}
	return val
}

func MaybeEnv(key string, dfault string) string {
	val := os.Getenv(key)
	if val == "" {
		return dfault
	}
	return val
}
