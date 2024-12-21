package application

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"net/http"

	"net/http/httptest"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		request        Request
		expected       string
		expectedStatus int
	}{
		{
			name:           "Valid Expression",
			request:        Request{Expression: "2+2*2"},
			expected:       "{\n  \"result\": 6.000000\n}",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Expression",
			request:        Request{Expression: "invalid"},
			expected:       "{\n  \"error\": \"invalid expression\"\n}",
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "Hello World Feature",
			request:        Request{Expression: "Hello world!"},
			expected:       "{\n  \"error\": \"It's not a bug. It's a feature\"\n}",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest("GET", "/api/v1/calculate", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			CalcHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()
			if err != nil {
				panic(err)
			}

			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, w.Code)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != tc.expected {
				t.Errorf("Expected response body \"%s\", but got \"%s\"", tc.expected, string(body))
			}
		})
	}
}
