package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"zaratedev/internal/models"

	"github.com/dlclark/regexp2"
)

var users map[int]models.User

type UserHandler struct{}

// Register maneja las solicitudes para registrar un nuevo usuario
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := register(newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("User registration: %+v\n", user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func register(user models.User) (models.User, error) {
	if user.Username == "" {
		return models.User{}, errors.New("username is requried")
	}

	if user.Email == "" {
		return models.User{}, errors.New("email is requried")
	}
	if user.Phone == "" {
		return models.User{}, errors.New("phone is requried")
	}

	if user.Password == "" {
		return models.User{}, errors.New("password is requried")
	}

	// Valid email
	if !validEmail(user.Email) {
		return models.User{}, errors.New("email not valid")
	}

	// Valid phone
	if !validPhone(user.Phone) {
		return models.User{}, errors.New("phone not valid")
	}

	// Valid password
	if !validPassword(user.Password) {
		return models.User{}, errors.New("password not valid")
	}

	// Exist email ?
	if existEmail(user.Email) {
		return models.User{}, errors.New("the email already exists")
	}

	// Exist phone?
	if existPhone(user.Phone) {
		return models.User{}, errors.New("the phone already exists")
	}

	users[len(users)+1] = user

	return user, nil
}

func validEmail(email string) bool {
	r := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + `{|}~-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return r.MatchString(email)
}

func validPhone(phone string) bool {
	r := regexp.MustCompile(`^\d{10}$`)
	return r.MatchString(phone)
}

func validPassword(password string) bool {
	// Regex
	re, _ := regexp2.Compile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[@$&])(?=.*[0-9])[a-zA-Z@$&0-9]{6,12}$`, 0)

	match, _ := re.MatchString(password)

	return match
}

func existEmail(email string) bool {
	return findUserByEmail(email)
}

func existPhone(phone string) bool {
	return findUserByPhone(phone)
}

func findUserByEmail(email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}

	return false
}

func findUserByPhone(phone string) bool {
	for _, user := range users {
		if user.Phone == phone {
			return true
		}
	}

	return false
}

// NewUserHandler crea una nueva instancia de UserHandler
func NewUserHandler(userMap map[int]models.User) *UserHandler {
	users = userMap
	return &UserHandler{}
}
