package handlers

import (
    "github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
    bookHandler := NewBookHandler()
    
    router.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
}
