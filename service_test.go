package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a App

func TestGenerateHandler(t *testing.T) {
	testServiceCases := []struct {
		route string
		want  int
		name  string
	}{
		{"5", 5, "Standard parentheses string"},
		{"1000", 1000, "Long parentheses string"},
		{"1", 1, "Short parentheses string"},
		{"0", 0, "Emptyy parentheses string"},
	}

	for _, tc := range testServiceCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/generate/"+tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()

			a.New()
			a.Router.ServeHTTP(response, req)
			assert.Equal(t, http.StatusOK, response.Code, "Expected response code %d. Got %d\n", http.StatusOK, response.Code)
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
