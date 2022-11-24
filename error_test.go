package respond

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithError(t *testing.T) {
	var tests = []struct {
		msg    string
		status int
	}{
		{
			msg:    "not found",
			status: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			// Arrange.
			w := httptest.NewRecorder()

			WithError(w, tt.msg, tt.status)

			// Assert.
			expected := map[string]interface{}{
				"message":    tt.msg,
				"statusCode": float64(tt.status),
			}
			if tt.status != w.Result().StatusCode {
				t.Errorf("expected status: %d, got %d", tt.status, w.Result().StatusCode)
			}
			actual := map[string]interface{}{}
			dec := json.NewDecoder(w.Result().Body)
			_ = dec.Decode(&actual)
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}

