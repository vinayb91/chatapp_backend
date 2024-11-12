package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vinayb91/chatapp_backend/config"
	"github.com/vinayb91/chatapp_backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB() // Establish MongoDB connection

	r := mux.NewRouter()
	routes.AuthRoutes(r) // Setup auth routes

	port := ":5000"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
