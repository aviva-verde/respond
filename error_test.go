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

func TestWithBadRequest(t *testing.T) {
	w := httptest.NewRecorder()

	WithBadRequest(w, "Test message")
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode, "status code does not match")
	httpassert.JSONBodyEqual(t, Error{
		Message:    "Test message",
		StatusCode: http.StatusBadRequest,
	}, w.Result().Body)
}

func TestWithUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()

	WithUnauthorized(w, "Test message")
	assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode, "status code does not match")
	httpassert.JSONBodyEqual(t, Error{
		Message:    "Test message",
		StatusCode: http.StatusUnauthorized,
	}, w.Result().Body)
}

func TestWithForbidden(t *testing.T) {
	w := httptest.NewRecorder()

	WithForbidden(w, "Test message")
	assert.Equal(t, http.StatusForbidden, w.Result().StatusCode, "status code does not match")
	httpassert.JSONBodyEqual(t, Error{
		Message:    "Test message",
		StatusCode: http.StatusForbidden,
	}, w.Result().Body)
}

func TestWithNotFond(t *testing.T) {
	w := httptest.NewRecorder()

	WithNotFound(w, "Test message")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode, "status code does not match")
	httpassert.JSONBodyEqual(t, Error{
		Message:    "Test message",
		StatusCode: http.StatusNotFound,
	}, w.Result().Body)
}

func TestWithMethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()

	WithMethodNotAllowed(w, "Test message")
	assert.Equal(t, http.StatusMethodNotAllowed, w.Result().StatusCode, "status code does not match")
	httpassert.JSONBodyEqual(t, Error{
		Message:    "Test message",
		StatusCode: http.StatusMethodNotAllowed,
	}, w.Result().Body)
}

func TestWithInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()

	WithInternalServerError(w, "Test message")
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode, "status code does not match")
	httpassert.JSONBodyEqual(t, Error{
		Message:    "Test message",
		StatusCode: http.StatusInternalServerError,
	}, w.Result().Body)
}
