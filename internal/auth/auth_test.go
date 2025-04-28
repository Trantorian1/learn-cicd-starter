package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test valid API key
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey test-key-123")
	
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if key != "test-key-123" {
		t.Errorf("Expected key 'test-key-123', got '%s'", key)
	}
	
	// Test missing header
	emptyHeaders := http.Header{}
	_, err = GetAPIKey(emptyHeaders)
	if err == nil {
		t.Error("Expected error for missing header, got nil")
	}
}
