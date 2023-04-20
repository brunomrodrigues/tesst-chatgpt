package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConvertHandler(t *testing.T) {
	// Prepare test cases
	testCases := []struct {
		brl       string
		currency  string
		status    int
		resp      map[string]interface{}
		errMsg    string
	}{
		// Valid test cases
		{
			brl:      "100",
			currency: "USD",
			status:   http.StatusOK,
			resp: map[string]interface{}{
				"brl":       100.0,
				"currency": "USD",
				"converted": 19.0,
			},
			errMsg: "",
		},
		{
			brl:      "1500",
			currency: "EUR",
			status:   http.StatusOK,
			resp: map[string]interface{}{
				"brl":       1500.0,
				"currency": "EUR",
				"converted": 242.25,
			},
			errMsg: "",
		},
		// Invalid test cases
		{
			brl:      "invalid",
			currency: "USD",
			status:   http.StatusBadRequest,
			resp:     nil,
			errMsg:   "Invalid BRL value",
		},
		{
			brl:      "100",
			currency: "invalid",
			status:   http.StatusBadRequest,
			resp:     nil,
			errMsg:   "Invalid currency",
		},
	}

	// Test each case
	for _, tc := range testCases {
		// Create request
		req, err := http.NewRequest("GET", "/convert", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.URL.RawQuery = "brl=" + tc.brl + "&currency=" + tc.currency

		// Create response recorder
		rec := httptest.NewRecorder()

		// Call handler
		convertHandler(rec, req)

		// Check response status code
		if rec.Code != tc.status {
			t.Errorf("Expected status code %d but got %d", tc.status, rec.Code)
		}

		// Check response body and error message
		if tc.resp != nil {
			var resp map[string]interface{}
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
				t.Errorf("Failed to decode response body: %v", err)
			} else if resp["brl"] != tc.resp["brl"] || resp["currency"] != tc.resp["currency"] || resp["converted"] != tc.resp["converted"] {
				t.Errorf("Invalid response body, expected %+v but got %+v", tc.resp, resp)
			}
		} else if !strings.Contains(rec.Body.String(), tc.errMsg) {
			t.Errorf("Invalid error message, expected %q but got %q", tc.errMsg, rec.Body.String())
		}
	}
}
