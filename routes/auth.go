package routes

import (
	"github.com/abahernest/greed-api/controller"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router){
	r.HandleFunc("/signup", controller.Signup).Methods("GET")
	r.HandleFunc("/signin", controller.Signin).Methods("GET")
}