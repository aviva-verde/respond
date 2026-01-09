package respond

import (
	"encoding/json"
	"net/http"
)

// WithJSON responds with the supplies status code and response body, v. The body will be encoded
// as JSON. Note that, if v is nil, the body will not be blank but will instead be "null".
func WithJSON(w http.ResponseWriter, v any, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		WithError(w, "failed to encode", http.StatusInternalServerError)
	}
}

// WithStatusOKAndJSON responds with a 200 (OK) response code and the supplied response body. The
// body will be encoded as JSON. Note that, if v is nil, the body will not be blank but will instead
// be "null".
func WithStatusOKAndJSON(w http.ResponseWriter, v any) {
	WithJSON(w, v, http.StatusOK)
}

// WithStatusCreatedAndJSON responds with a 201 (Created) response code and the supplied response
// body. The body will be encoded as JSON. Note that, if v is nil, the body will not be blank but
// will instead be "null".
func WithStatusCreatedAndJSON(w http.ResponseWriter, v any) {
	WithJSON(w, v, http.StatusCreated)
}

// WithStatusAcceptedAndJSON responds with a 202 (Accepted) response code and the supplied response
// body. The body will be encoded as JSON. Note that, if v is nil, the body will not be blank but
// will instead be "null".
func WithStatusAcceptedAndJSON(w http.ResponseWriter, v any) {
	WithJSON(w, v, http.StatusAccepted)
}
