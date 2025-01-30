package storage

import (
	"fetch_backend/models"
	"strconv"
	"sync"
	"time"
)

var (
	receiptMap = make(map[string]models.Receipt)
	mu         sync.Mutex
)

// StoreReceipt saves a receipt and returns its generated ID.
func StoreReceipt(receipt models.Receipt) string {
	mu.Lock()
	id := GenerateReceiptID()
	receiptMap[id] = receipt
	defer mu.Unlock()
	return id
}

// GetReceipt retrieves a receipt by ID.
func GetReceipt(id string) (models.Receipt, bool) {
	mu.Lock()
	defer mu.Unlock()
	receipt, exists := receiptMap[id]
	return receipt, exists
}

// GenerateReceiptID returns a unique string ID.
func GenerateReceiptID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
