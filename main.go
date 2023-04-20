package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ExchangeRate struct {
	Result   string             `json:"result"`
	From     string             `json:"from"`
	To       string             `json:"to"`
	Exchange string             `json:"exchange_rate"`
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse request parameters", http.StatusBadRequest)
		return
	}

	// Get input values
	brlStr := r.Form.Get("brl")
	currency := r.Form.Get("currency")

	// Convert BRL to float
	brl, err := strconv.ParseFloat(brlStr, 64)
	if err != nil {
		http.Error(w, "Invalid BRL value", http.StatusBadRequest)
		return
	}

	// Fetch exchange rate from API
	resp, err := http.Get(fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/BRL"))
	if err != nil {
		http.Error(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Parse exchange rate response
	var exchangeRate map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRate); err != nil {
		http.Error(w, "Failed to parse exchange rate", http.StatusInternalServerError)
		return
	}

	// Extract exchange rate for target currency
	rates, ok := exchangeRate["rates"].(map[string]interface{})
	if !ok {
		http.Error(w, "Failed to extract exchange rates", http.StatusInternalServerError)
		return
	}
	rate, ok := rates[currency].(float64)
	if !ok {
		http.Error(w, "Invalid currency", http.StatusBadRequest)
		return
	}

	// Calculate converted value
	converted := brl * rate

	// Write response
	response := map[string]interface{}{
		"brl":       brl,
		"currency":  currency,
		"converted": converted,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/convert", convertHandler)
	http.ListenAndServe(":8080", nil)
}
