package main

import (
	"fetch_backend/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on :8080...")

	http.HandleFunc("/receipts/process", handlers.ProcessReceipt)
	http.HandleFunc("/receipts/", handlers.GetPoints)

	http.ListenAndServe(":8080", nil)
}
