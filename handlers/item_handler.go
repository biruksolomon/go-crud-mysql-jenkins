package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/biruksolomon/go-crud-portfolio/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var items []models.Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	item.ID = uuid.New().String()
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedItem models.Item
	json.NewDecoder(r.Body).Decode(&updatedItem)
	for i, item := range items {
		if item.ID == params["id"] {
			updatedItem.ID = item.ID
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range items {
		if item.ID == params["id"] {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
