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

// CheckErr : helper func to check if an error
func CheckErr(err error) {
	if err != nil {
		panic(err) // panic stops ordinary flow of control and begins panicking (program crashes)
	}
}
