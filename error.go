package respond

import (
	"encoding/json"
	"net/http"
)

// Error is a standard error response.
type Error struct {
	// Message is the error message.
	Message string `json:"message" example:"The request was invalid."`
	// StatusCode is the HTTP status code.
	StatusCode int `json:"statusCode" example:"400"`
	// Issues is a list of issues with the request. This is optional.
	Issues []string `json:"issues,omitempty" example:"The request was invalid."`
}

// WithError an error response to the response writer with given status code. This includes an error
// message, msg, and an optional list of issues. The resulting body will be an Error struct
// encoded as JSON.
func WithError(w http.ResponseWriter, msg string, status int, issues ...string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	v := Error{
		Message:    msg,
		StatusCode: status,
		Issues:     issues,
	}
	_ = json.NewEncoder(w).Encode(v)
}
