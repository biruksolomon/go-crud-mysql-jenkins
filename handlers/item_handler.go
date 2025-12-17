package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/biruksolomon/go-crud-portfolio/config"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	_ "github.com/biruksolomon/go-crud-portfolio/config"
	"github.com/biruksolomon/go-crud-portfolio/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query("SELECT id, name, price FROM items")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.Name, &item.Price)
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var item models.Item
	query := "SELECT id, name, price FROM items WHERE id=?"
	err := config.DB.QueryRow(query, params["id"]).Scan(
		&item.ID, &item.Name, &item.Price,
	)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	item.ID = uuid.New().String()

	query := "INSERT INTO items (id, name, price) VALUES (?, ?, ?)"
	_, err := config.DB.Exec(query, item.ID, item.Name, item.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	query := "UPDATE items SET name=?, price=? WHERE id=?"
	result, err := config.DB.Exec(query, item.Name, item.Price, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	item.ID = params["id"]
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query := "DELETE FROM items WHERE id=?"
	result, err := config.DB.Exec(query, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
