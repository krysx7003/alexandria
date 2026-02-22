package main

import (
	"log"
	"os"
	"net/http"

    "github.com/gorilla/mux"
	"alexandria/internal/handlers"
)

func main(){
	router := mux.NewRouter()

	handlers.SetupRoutes(router)

	router.Use(loggingMiddleware)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
	}

	http.ListenAndServe(":"+port, router)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
    })
}
