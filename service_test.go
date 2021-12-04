package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHandler(t *testing.T) {
	testServiceCases := []struct {
		length string
		want   int
		name   string
		status int
	}{
		{"5", 5, "Standard value", http.StatusOK},
		{"1000", 1000, "Long value", http.StatusOK},
		{"1", 1, "Short value", http.StatusOK},
		{"0", 0, "Zero value", http.StatusOK},
		{"?", 0, "Non int char", http.StatusBadRequest},
		{"-1", 0, "Negative value", http.StatusBadRequest},
	}

	for _, tc := range testServiceCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/generate/"+tc.length, nil)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			generate(response, req)
			assert.Equal(t, tc.status, response.Code, "Expected response code %d. Got %d\n", tc.status, response.Code)

			if tc.status != http.StatusOK {
				t.Skip()
			}

			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			var lines Parentheses
			err = json.Unmarshal(data, &lines)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			assert.Equal(t, tc.want, len(lines.Line), "Expected line length %d. Got %d\n", tc.want, len(lines.Line))
		})
	}
}
