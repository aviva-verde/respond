package respond

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aviva-verde/respond/httpassert"
	"github.com/stretchr/testify/assert"
)

func TestEmptyBody(t *testing.T) {
	var tests = []struct {
		name   string
		status int
	}{
		{
			name:   http.StatusText(http.StatusOK),
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			w := httptest.NewRecorder()

			WithEmptyBody(w, tt.status)
			assert.Equal(t, tt.status, w.Result().StatusCode, "status code does not match")
			httpassert.JSONBodyEqual(t, nil, w.Result().Body)
		})
	}
}
