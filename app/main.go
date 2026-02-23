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

	// myString := "Hello World!"
	//
	// bytes := []byte(myString)
	// err := os.WriteFile("/data/myfile.txt", bytes, 0666)
	//
	// if err != nil {
	// 	panic("Failed to write file")
	// }
	//
	// log.Printf("File created successfully!")
	http.ListenAndServe(":"+port, router)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
    })
}
