package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKeyHappy(t *testing.T) {
	testApiKey := "only-a-test-api-key"
	testHeader := make(http.Header)
	testHeader.Add("Authorization", fmt.Sprintf("ApiKey %s", testApiKey))

	result, err := GetAPIKey(testHeader)
	if err != nil {
		t.Errorf("Happy path test failed with error: %v", err)
	}

	if result != testApiKey {
		t.Errorf("Happy path test failed--did not return the correct key from the headers:\nexpected: %s\nresult: %s", testApiKey, result)
	}
}

func TestGetAPIKeyEmptyKey(t *testing.T) {
	// testApiKey := ""
	testHeader := make(http.Header)
	testHeader.Add("Authorization", "")

	_, err := GetAPIKey(testHeader)
	if err == nil {
		t.Errorf("Empty API key sad path test failed: should have returned an error")
	}
}

func TestGetAPIKeyEmptyAuthHeader(t *testing.T) {
	testHeader := make(http.Header)

	_, err := GetAPIKey(testHeader)
	if err == nil {
		t.Errorf("Empty Authorization header sad path test failed: should have returned an error")
	}
}
