package api

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todo", TodoList).Methods("GET")
	router.HandleFunc("/todo", TodoCreate).Methods("POST")
	router.HandleFunc("/todo/{id:[0-9]+}", TodoGet).Methods("GET")
	router.HandleFunc("/todo/{id:[0-9]+}", TodoUpdate).Methods("PUT")
	router.HandleFunc("/todo/{id:[0-9]+}", TodoDelete).Methods("DELETE")

	return router
}
