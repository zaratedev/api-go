package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"zaratedev/internal/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// AuthHandler maneja las operaciones relacionadas con la autenticación
type AuthHandler struct{}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login maneja las solicitudes de inicio de sesión
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if creds.Username == "" && creds.Email == "" {
		http.Error(w, "Username or email is required", http.StatusBadRequest)
		return
	}

	if creds.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	var user *models.User
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

// NewAuthHandler crea una nueva instancia de AuthHandler
func NewAuthHandler(userMap map[int]models.User) *AuthHandler {
	users = userMap
	return &AuthHandler{}
}
