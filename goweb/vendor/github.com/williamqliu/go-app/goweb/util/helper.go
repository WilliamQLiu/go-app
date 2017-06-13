package util

import (
	"encoding/json"
	"net/http"
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

// RespondWithError : respond with map of "error" and "errror message"
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON : respond with JSON payload
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
