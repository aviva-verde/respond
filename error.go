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

// WithBadRequest creates an error response with a status code of 400 (Bad Request). This includes
// an error message, msg, and an optional list of issues. The resulting body will be an Error struct
// encoded as JSON.
func WithBadRequest(w http.ResponseWriter, msg string, issues ...string) {
	WithError(w, msg, http.StatusBadRequest, issues...)
}

// WithUnauthorized creates an error response with a status code of 401 (Unauthorized). This includes
// an error message, msg, and an optional list of issues. The resulting body will be an Error struct
// encoded as JSON.
func WithUnauthorized(w http.ResponseWriter, msg string, issues ...string) {
	WithError(w, msg, http.StatusUnauthorized, issues...)
}

// WithForbidden creates an error response with a status code of 403 (Forbidden). This includes an
// error message, msg, and an optional list of issues. The resulting body will be an Error struct
// encoded as JSON.
func WithForbidden(w http.ResponseWriter, msg string, issues ...string) {
	WithError(w, msg, http.StatusForbidden, issues...)
}

// WithNotFound creates an error response with a status code of 404 (Not Found). This includes an
// error message, msg, and an optional list of issues. The resulting body will be an Error struct
// encoded as JSON.
func WithNotFound(w http.ResponseWriter, msg string, issues ...string) {
	WithError(w, msg, http.StatusNotFound, issues...)
}

// WithMethodNotAllowed creates an error response with a status code of 405 (Method Not Allowed).
// This includes an error message, msg, and an optional list of issues. The resulting body will be
// an Error struct encoded as JSON.
func WithMethodNotAllowed(w http.ResponseWriter, msg string, issues ...string) {
	WithError(w, msg, http.StatusMethodNotAllowed, issues...)
}

// WithInternalServerError creates an error response with a status code of 500 (Internal Server
// Error). This includes an error message, msg, and an optional list of issues. The resulting body will be
// an Error struct encoded as JSON.
func WithInternalServerError(w http.ResponseWriter, msg string, issues ...string) {
	WithError(w, msg, http.StatusInternalServerError, issues...)
}
