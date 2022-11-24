package respond

import (
	"encoding/json"
	"net/http"
)

func WithError(w http.ResponseWriter, msg string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	v := map[string]any{
		"message":    msg,
		"statusCode": int64(status),
	}
	json.NewEncoder(w).Encode(v)
}
