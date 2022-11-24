package respond

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			var actual Error
			json.NewDecoder(w.Result().Body).Decode(&actual)
			if diff := cmp.Diff(tt, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
