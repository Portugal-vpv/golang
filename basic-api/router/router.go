package router

import (
	"api/api/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("DELETE", "OPTIONS")
	return router
}
