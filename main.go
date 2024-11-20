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

	http.HandleFunc("/items", handlers.GetItems)

	log.Println("Server running at http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
