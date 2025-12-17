package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/biruksolomon/go-crud-portfolio/config"
	"github.com/biruksolomon/go-crud-portfolio/handlers"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	log.Println("🚀 Server running on :8381")
	log.Fatal(http.ListenAndServe(":8381", router))
}
