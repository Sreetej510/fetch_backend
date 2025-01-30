package utils

import (
	"fetch_backend/models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "M&M Corner Market receipt",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Total:        "9.00",
				Items: []models.Item{
					{"Gatorade", "2.25"},
					{"Gatorade", "2.25"},
					{"Gatorade", "2.25"},
					{"Gatorade", "2.25"},
				},
			},
			expected: 109,
		},
		{
			name: "Target receipt",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Total:        "35.35",
				Items: []models.Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
			},
			expected: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculatePoints(tt.receipt)
			if err != nil {
				t.Errorf("Error while calculating points %s", err.Error())
			}
			if got != tt.expected {
				t.Errorf("CalculatePoints() = %d, expected %d", got, tt.expected)
			}
		})
	}
}
