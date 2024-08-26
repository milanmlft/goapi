package api

import (
	"encoding/json"
	"net/http"
)

type IceCreamParams struct {
	Username string
}

// Ice Cream API response
type IceCreamResponse struct {
	Flavours []string
	Code     int
}

// Error response
type Error struct {
	Message string
	Code    int
}

// Function to write errors to the HTTP response writer
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	// To return specific errors in response, useful for user
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// To return a generic error message, e.g. internal bugs, not useful for user
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
