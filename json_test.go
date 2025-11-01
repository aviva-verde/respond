package respond

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aviva-verde/respond/httpassert"
	"github.com/stretchr/testify/assert"
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

			WithJSON(w, tt.item, http.StatusOK)

			// Assert.
			assert.Equal(t, http.StatusOK, w.Result().StatusCode, "status code does not match")
			httpassert.JSONBodyEqual(t, tt.expectedBody, w.Result().Body)
		})
	}
}
