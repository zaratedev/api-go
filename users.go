package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/dlclark/regexp2"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

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

func register(user User) (User, error) {
	if user.Username == "" {
		return User{}, errors.New("username is requried")
	}

	if user.Email == "" {
		return User{}, errors.New("email is requried")
	}
	if user.Phone == "" {
		return User{}, errors.New("phone is requried")
	}

	if user.Password == "" {
		return User{}, errors.New("password is requried")
	}

	// Valid email
	if !validEmail(user.Email) {
		return User{}, errors.New("email not valid")
	}

	// Valid phone
	if !validPhone(user.Phone) {
		return User{}, errors.New("phone not valid")
	}

	// Valid password
	if !validPassword(user.Password) {
		return User{}, errors.New("password not valid")
	}

	// Exist email ?
	if existEmail(user.Email) {
		return User{}, errors.New("the email already exists")
	}

	// Exist phone?
	if existPhone(user.Phone) {
		return User{}, errors.New("the phone already exists")
	}

	id++
	users[id] = user

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
