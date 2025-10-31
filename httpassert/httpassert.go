package httpassert

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// JSONBodyEqual compares if the body of an HTTP response is expected. Usage:
//
//	expected := Model{ Message: "Hello, world!" } // the expected body
//	httpassert.JSONBodyEqual(t, expected, response.Body)
func JSONBodyEqual(t *testing.T, expected interface{}, actual io.Reader) {
	if actual == nil {
		if expected == nil {
			return
		}
		t.Error("expected a body but received nil")
	}

	bodyBytes, err := io.ReadAll(actual)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	if expected == nil {
		assert.Empty(t, bodyBytes, "expected no body but one was received")
	}

	var expectedBody string
	switch value := expected.(type) {
	case string:
		expectedBody = value
	default:
		expectedBodyBytes, err := json.Marshal(expected)
		if err != nil {
			t.Fatalf("could not marshal expected response body: %v", err)
		}
		expectedBody = string(expectedBodyBytes)
	}

	assert.Equal(
		t,
		strings.TrimSpace(expectedBody),
		strings.TrimSpace(string(bodyBytes)),
	)
}
