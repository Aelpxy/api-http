package main

import (
	"log"
	"net/http"

	"github.com/aelpxy/api-http/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/u", handlers.GetAllUsers).Methods(http.MethodGet)

	router.HandleFunc("/u/{id}", handlers.GetUser).Methods(http.MethodGet)

	router.HandleFunc("/u", handlers.PostUser).Methods(http.MethodPost)

	router.HandleFunc("/u/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	router.HandleFunc("/u/{id}", handlers.PutUser).Methods(http.MethodPut)

	log.Println("Listening on PORT 8080")
	http.ListenAndServe(":8080", router)
}
