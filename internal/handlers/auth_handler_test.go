package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zaratedev/internal/models"
)

func TestLoginUser(t *testing.T) {
	users = make(map[int]models.User)
	users = map[int]models.User{
		1: {Username: "user1", Email: "user1@example.com", Phone: "1234567890", Password: "password1"},
		2: {Username: "user2", Email: "user2@example.com", Phone: "9876543210", Password: "password2"},
	}

	validCredentials := models.Credentials{
		Username: "user1",
		Email:    "user1@example.com",
		Password: "password1",
	}

	reqBody, _ := json.Marshal(validCredentials)
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := NewAuthHandler(users)
	handler.Login(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}

	var responseMap map[string]string
	err := json.NewDecoder(rr.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	_, ok := responseMap["token"]
	if !ok {
		t.Error("Token not found in response")
	}
}
