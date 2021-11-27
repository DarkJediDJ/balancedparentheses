package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/generate/666", nil)
	if err != nil {
		fmt.Println("alkfk") // HERE - ?
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	a.New()
	a.Router.ServeHTTP(response, req)

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}

	// HERE - do check here
	// HERE - use test struct with cases
}
