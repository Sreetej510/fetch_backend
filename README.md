# Receipt Processor API (FETCH REWARDS BACKEND)

This is a RESTful API that processes receipt data and calculates reward points based on predefined rules from [fetch rewards](https://github.com/fetch-rewards/receipt-processor-challenge/)

## Features  
- Accepts receipt data and processes it.  
- Calculates reward points based on purchase details.  
- Retrieves stored receipt points.  

## Tech Stack  
- **Language:** Go  
- **Framework:** net/http  
- **Database:** No database (as per request) 

---

## Setup Instructions  

### Prerequisites  
- **Go 1.18+**
- **cURL or Postman** for API testing

### Run the Application  
```sh
go run main.go
```

The server should now be running on **`http://localhost:8080`**.  

---

## API Endpoints  

### **1. Process Receipt**  
_Submit a receipt for processing._  

**Endpoint:**  
```http
POST /receipts/process
```
**Request Body (JSON):**  
```json
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
```
**Response:**  
```json
{ "id": "unique-id (nano second as string)" }
```

---

### **2. Get Points**  
_Retrieve points for a specific receipt ID._  

**Endpoint:**  
```http
GET /receipts/{id}/points
```
**Response:**  
```json
{ "points": 10 }
```

---

## Running Tests  
```sh
go test ./utils -v
```

---
