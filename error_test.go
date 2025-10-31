package respond

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aviva-verde/respond/httpassert"
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
			if tt.StatusCode != w.Result().StatusCode {
				t.Errorf("expected status: %d, got %d", tt.StatusCode, w.Result().StatusCode)
			}
			httpassert.JSONBodyEqual(t, tt, w.Result().Body)
		})
	}
}
