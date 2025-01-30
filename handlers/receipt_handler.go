package handlers

import (
	"encoding/json"
	"fetch_backend/models"
	"fetch_backend/storage"
	"fetch_backend/utils"
	"net/http"
	"strings"
)

// ProcessReceipt handles POST /receipts/process
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate ID & store receipt
	receiptID := storage.StoreReceipt(receipt)

	// Respond with receipt ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": receiptID})
}

// GetPoints handles GET /receipts/{id}
func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")

	receipt, exists := storage.GetReceipt(id)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Calculate points
	points, err := utils.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
