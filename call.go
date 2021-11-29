package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

//Call requests service for 1000 times and finds % of balanced lines
func Call() map[int]int {
	res := make(map[int]int)
	const requests int = 1000
	for j := 2; j <= 8; j *= 2 {
		count := 0
		for i := 1; i <= requests; i++ {
			req, err := http.NewRequest("GET", "/generate/"+strconv.Itoa(j), nil)
			if err != nil {
				log.Fatal(err)
			}

			response := httptest.NewRecorder()

			A.New()
			A.Router.ServeHTTP(response, req)

			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var lines Parentheses

			err = json.Unmarshal(data, &lines)
			if err != nil {
				log.Fatal(err)
			}

			if Balancer(lines.Line) {
				count++
			}
		}
		res[j] = ((count / requests) * 100)
	}
	return res
}
