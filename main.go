package main

import (
	"log"
	"net/http"

	"github.com/biruksolomon/go-crud-portfolio/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
