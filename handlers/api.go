package handlers

import (
	"encoding/json"
	"maxchat/models"
	"net/http"
	"strings"
)

var items []models.Item

func InitData(data []models.Item) {
	items = data
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	modelFilter := query.Get("model")
	techFilter := query.Get("tech")

	filtered := items
	if modelFilter != "" {
		filtered = filterByModel(filtered, modelFilter)
	}
	if techFilter != "" {
		filtered = filterByTech(filtered, techFilter)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for _, item := range items {
		if item.Code == newItem.Code {
			http.Error(w, "Item with the same code already exists", http.StatusConflict)
			return
		}
	}

	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	var updatedItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.Code == code {
			items[i] = updatedItem
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.Code == code {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func filterByModel(data []models.Item, model string) []models.Item {
	var result []models.Item
	for _, item := range data {
		if item.Model == model {
			result = append(result, item)
		}
	}
	return result
}

func filterByTech(data []models.Item, tech string) []models.Item {
	var result []models.Item
	for _, item := range data {
		for _, t := range item.Tech {
			if strings.Contains(t, tech) {
				result = append(result, item)
				break
			}
		}
	}
	return result
}
