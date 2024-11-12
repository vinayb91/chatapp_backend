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
	config.ConnectDB()

	r := mux.NewRouter()
	routes.AuthRoutes(r)

	port := ":5000"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
