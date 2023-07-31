package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-token-auth-jwt-mux/config"
	"github.com/mhdianrush/go-token-auth-jwt-mux/controllers"
	"github.com/mhdianrush/go-token-auth-jwt-mux/middlewares"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()

	routes := mux.NewRouter()

	routes.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	routes.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)
	routes.HandleFunc("/logout", controllers.Logout).Methods(http.MethodGet)

	// this route will be protect with middleware
	// if user isn't login, it will return an unauthorized response.
	api := routes.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", controllers.Index).Methods(http.MethodGet)

	// use middleware
	api.Use(middlewares.JWTMiddleware)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Printf("failed create log file %s", err.Error())
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: routes,
	}
	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}

	logger.Printf("server running on port %s", os.Getenv("SERVER_PORT"))
}
