package routes

import (
	"github.com/gorilla/mux"
)

func  SetupRoutes(r *mux.Router) {
	// auth router
	authRouter := r.PathPrefix("/auth").Subrouter()
	AuthRoutes(authRouter)
	
	// user router

}