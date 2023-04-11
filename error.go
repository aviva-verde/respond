package respond

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"statusCode"`
	Issues     []string `json:"issues,omitempty"`
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
