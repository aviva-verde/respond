package respond

import "net/http"

// WithEmptyBody writes the status code to the response writer and keeps the response body empty.
func WithEmptyBody(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
