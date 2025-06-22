package main

import (
	"golang-watchlist/internal/db"
	"golang-watchlist/internal/initializers"
	"golang-watchlist/internal/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	router := mux.NewRouter()

	routes.HandleRecordRoutes(router, database)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port :%v\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
