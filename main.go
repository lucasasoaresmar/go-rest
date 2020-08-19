package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"example.com/auth/adapters/database"
	"example.com/auth/adapters/token"
	"example.com/auth/controllers"
)

const port = ":3000"

func main() {
	routes()
}

func routes() {
	router := mux.NewRouter()
	v1Router := router.PathPrefix("/v1").Subrouter()

	pingHandler := controllers.Ping{}
	v1Router.HandleFunc("/ping", pingHandler.Pong).Methods("GET")

	userRepo := database.User{}
	tokenLib := token.MockedToken{ExpectedToken: "Mocked Token"}

	authHandler := controllers.Auth{
		UserRepository: &userRepo,
		Token:          &tokenLib,
	}
	v1Router.HandleFunc("/login", authHandler.Login).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(port, router)
}
