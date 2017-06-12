package util

import (
	"os"
)

// GetKey : helper func to get a 'key', if none then returns 'fallback' string
func GetKey(key, fallback string) string {
	value, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	return value
}
