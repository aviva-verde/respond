package respond

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aviva-verde/respond/httpassert"
	"github.com/stretchr/testify/assert"
)

func TestWithError(t *testing.T) {
	var tests = []Error{
		{
			Message:    "not found",
			StatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.Message, func(t *testing.T) {
			// Arrange.
			w := httptest.NewRecorder()

			WithError(w, tt.Message, tt.StatusCode)

			// Assert.
			assert.Equal(t, tt.StatusCode, w.Result().StatusCode, "status code does not match")
			httpassert.JSONBodyEqual(t, tt, w.Result().Body)
		})
	}
}
