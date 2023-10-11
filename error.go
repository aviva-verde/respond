package respond

import (
	"encoding/json"
	"net/http"
)

// Error is a standard error response.
type Error struct {
	// Message is the error message.
	// Example: "The request was invalid."
	Message string `json:"message"`
	// StatusCode is the HTTP status code.
	// Example: 400
	StatusCode int `json:"statusCode"`
	// Issues is a list of issues with the request. This is optional.
	// Example: ["The request was invalid."]
	Issues []string `json:"issues,omitempty"`
}

func WithError(w http.ResponseWriter, msg string, status int, issues ...string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	v := Error{
		Message:    msg,
		StatusCode: status,
		Issues:     issues,
	}
	json.NewEncoder(w).Encode(v)
}
