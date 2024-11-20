package main

import (
	"log"
	"maxchat/handlers"
	"maxchat/utils"
	"net/http"
)

func main() {
	data, err := utils.LoadData("./data/data.txt")
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}
	handlers.InitData(data)

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetItems(w, r)
		case http.MethodPost:
			handlers.CreateItem(w, r)
		case http.MethodPut:
			handlers.UpdateItem(w, r)
		case http.MethodDelete:
			handlers.DeleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running at http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
