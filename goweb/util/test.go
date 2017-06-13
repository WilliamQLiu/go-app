package util

import (
	"testing"
)

// CheckResponseCode : check if expected response code is same as actual response code
func CheckResponseCode(t *testing.T, expected, actual int) {
	// Compares response code
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
