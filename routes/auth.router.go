package routes

import (
	"github.com/vinayb91/chatapp_backend/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRouter := r.PathPrefix("/api/auth").Subrouter()
	// authRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	authRouter.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	// authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")
}
