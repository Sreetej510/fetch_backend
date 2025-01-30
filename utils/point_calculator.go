package utils

import (
	"fetch_backend/models"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CalculatePoints computes reward points for a given receipt.
func CalculatePoints(receipt models.Receipt) (int, error) {
	points := 0
	totalValue := parseTotalValue(receipt.Total)

	// 1 point per alphanumeric character in the retailer name
	points += len(regexp.MustCompile(`\w`).FindAllString(receipt.Retailer, -1))

	if totalValue > 0 && int(totalValue*100)%100 == 0 {
		// 75 (50 for whole and 25 for divided by 0.25) points if total is a whole number
		points += 75
	} else if totalValue > 0 && int(totalValue*100)%25 == 0 {
		// 25 for divided by 0.25 but not whole number
		points += 25
	}

	// 5 points per every 2 items
	points += (len(receipt.Items) / 2) * 5

	// item description is a multiple of 3
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, nil
			}
			points += int(math.Ceil(price * 0.2)) // multiply the price by 0.2 and round up to the nearest integer
		}
	}

	// 6 points if the day in the purchase date is odd
	day, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	if day.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is between 2:00pm and 4:00pm
	t, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	if t.Hour() >= 14 && t.Hour() < 16 {
		points += 10
	}
	return points, nil
}

// Helper function to parse total as float64
func parseTotalValue(total string) float64 {
	var value float64
	fmt.Sscanf(total, "%f", &value)
	return value
}
