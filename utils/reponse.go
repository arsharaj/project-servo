// Package utils provides utility functions for HTTP response formatting.
package utils

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse defines the standard structure for successful JSON responses.
type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse defines the standard structure for error JSON responses.
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// RespondJSON writes a JSON-encoded success response with the given status, message, and data.
func RespondJSON(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// RespondError writes a JSON-encoded error response with the given status and message.
func RespondError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(ErrorResponse{
		Status:  "error",
		Message: message,
	})
}
