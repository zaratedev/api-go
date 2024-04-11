package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zaratedev/internal/models"
)

func TestCreateUser(t *testing.T) {
	users = make(map[int]models.User)
	newUser := models.User{
		Username: "zaratedev",
		Email:    "zaratedev@gmail.com",
		Phone:    "5555555555",
		Password: "Test@123",
	}
	reqBody, _ := json.Marshal(newUser)
	req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))

	// Crear un ResponseRecorder para capturar la respuesta
	rr := httptest.NewRecorder()

	userHandler := NewUserHandler(users)
	userHandler.Register(rr, req)

	// Verificar el código de estado de la respuesta
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}

	var responseUser models.User
	err := json.NewDecoder(rr.Body).Decode(&responseUser)
	if err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	if responseUser.Username != newUser.Username || responseUser.Email != newUser.Email || responseUser.Phone != newUser.Phone || responseUser.Password != newUser.Password {
		t.Errorf("Response does not match expected user data")
	}
}

func TestValidEmail(t *testing.T) {
	// email válido
	email := "test@example.com"
	if !validEmail(email) {
		t.Errorf("Expected email %s to be valid, but it was not", email)
	}

	// email inválido
	email = "invalid-email"
	if validEmail(email) {
		t.Errorf("Expected email %s to be invalid, but it was not", email)
	}
}

func TestValidPhone(t *testing.T) {
	// teléfono válido
	phone := "1234567890"
	if !validPhone(phone) {
		t.Errorf("Expected phone number %s to be valid, but it was not", phone)
	}

	// teléfono inválido
	phone = "123"
	if validPhone(phone) {
		t.Errorf("Expected phone number %s to be invalid, but it was not", phone)
	}
}

func TestValidPassword(t *testing.T) {
	password := "Secure&567"
	if !validPassword(password) {
		t.Errorf("Expected password %s to be valid, but it was not", password)
	}

	password = "password" // No contiene caracteres especiales ni mayúsculas
	if validPassword(password) {
		t.Errorf("Expected password %s to be invalid, but it was not", password)
	}
}

func TestExistEmail(t *testing.T) {
	users = map[int]models.User{
		1: {Username: "user1", Email: "user1@example.com", Phone: "1234567890", Password: "password1"},
		2: {Username: "user2", Email: "user2@example.com", Phone: "9876543210", Password: "password2"},
	}

	email := "user1@example.com"
	if !existEmail(email) {
		t.Errorf("Expected email %s to exist, but it was not found", email)
	}

	email = "nonexisting@example.com"
	if existEmail(email) {
		t.Errorf("Expected email %s to not exist, but it was found", email)
	}
}

func TestExistPhone(t *testing.T) {
	users = map[int]models.User{
		1: {Username: "user1", Email: "user1@example.com", Phone: "1234567890", Password: "password1"},
		2: {Username: "user2", Email: "user2@example.com", Phone: "9876543210", Password: "password2"},
	}

	phone := "1234567890"
	if !existPhone(phone) {
		t.Errorf("Expected phone number %s to exist, but it was not found", phone)
	}

	phone = "0000000000"
	if existPhone(phone) {
		t.Errorf("Expected phone number %s to not exist, but it was found", phone)
	}
}
