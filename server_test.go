package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {

	// Create new request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create new recorder to test against
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(helloWorld)
	handler.ServeHTTP(res, req)

	// Test for response status code
	if status := res.Code; status != http.StatusOK {
		t.Error("Excepted status code", http.StatusOK)
	}

	resposeContent := &Message{}
	if err := json.Unmarshal(res.Body.Bytes(), resposeContent); err != nil {
		t.Error("Invalid request json. Error:", err)
	}

	// Test for response content
	expected := "hello world"
	if resposeContent.Text != expected {
		t.Error("Expected response", expected)
	}
}
