package main

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// Representacion del request para el login
type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims representa los claims del token JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Función para autenticar usuario
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if creds.Username == "" || creds.Email == "" {
		http.Error(w, "Username or email is required", http.StatusBadRequest)
		return
	}

	if creds.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	var user *User
	for _, u := range users {
		if creds.Username == u.Username || creds.Email == u.Email {
			if creds.Password == u.Password {
				user = &u
				break
			} else {
				http.Error(w, "Username / password incorrect", http.StatusUnauthorized)
				return
			}
		}
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error to generate token JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
