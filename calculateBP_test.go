package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateBalancedPercentage(t *testing.T) {
	testCalculateBPCases := []struct {
		name   string
		status int
	}{
		{"Status test", http.StatusOK},
	}

	for _, tc := range testCalculateBPCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/generate/result", nil)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()

			A.New()
			A.Router.ServeHTTP(response, req)
			assert.Equal(t, tc.status, response.Code, "Expected response code %d. Got %d\n", tc.status, response.Code)

			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			assert.NotNil(t, data)
		})
	}
}
