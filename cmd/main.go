package main

import (
	"fmt"
	"log"
	"net/http"

	"zaratedev/internal/handlers"
	"zaratedev/internal/models"

	"github.com/gorilla/mux"
)

var id int
var users map[int]models.User

func main() {
	users = make(map[int]models.User)

	// Inicializar el enrutador
	router := mux.NewRouter()

	// Registro de manejadores
	registerHandlers(router)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerHandlers(router *mux.Router) {
	// Manejador para el registro de usuarios
	userHandler := handlers.NewUserHandler(users)
	router.HandleFunc("/api/users", userHandler.Register).Methods("POST")

	// Manejador para iniciar sesión de usuario
	authHandler := handlers.NewAuthHandler(users)
	router.HandleFunc("/api/login", authHandler.Login).Methods("POST")
}
