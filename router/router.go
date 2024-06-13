package router

import (
	"github.com/gorilla/mux"
	"github.com/usmantech001/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/createMovies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/deleteMovies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/updateMovies/{id}", controller.DeleteOneMovie).Methods("PUT")

	return router
}
