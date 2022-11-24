package respond

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithJSON(t *testing.T) {
	var tests = []struct {
		name         string
		item         any
		expectedBody string
	}{
		{
			name:         "nil",
			item:         nil,
			expectedBody: "null\n",
		},
		{
			name: "value",
			item: map[string]interface{}{
				"key": "value",
			},
			expectedBody: `{"key":"value"}` + "\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			w := httptest.NewRecorder()

			WithJSON(w, tt.item, 200)

			// Assert.
			if http.StatusOK != w.Result().StatusCode {
				t.Errorf("expected status 200, got %d", w.Result().StatusCode)
			}
			actualBody, err := io.ReadAll(w.Result().Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}
			if diff := cmp.Diff(tt.expectedBody, string(actualBody)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
