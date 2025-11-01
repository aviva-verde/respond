package respond

import (
	"encoding/json"
	"net/http"
)

// WithJSON sets will write the supplied status and body, v, to the response writer. The body will
// be encoded as JSON. Note that if v is nil, the body will not be blank but will instead be "null".
func WithJSON(w http.ResponseWriter, v any, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		WithError(w, "failed to encode", http.StatusInternalServerError)
	}
}
