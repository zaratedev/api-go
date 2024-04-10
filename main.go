package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Estructura global de user
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Mapa global para simular db
var id int
var users map[int]User

func main() {
	users = make(map[int]User)

	router := mux.NewRouter()

	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/login", LoginUser).Methods("POST")

	// Iniciar servidor en el puerto 8080
	fmt.Println("Server :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
