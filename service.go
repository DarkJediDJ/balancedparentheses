package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func generate(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.RequestURI()
	length := URL[strings.LastIndex(URL, "/")+1:]
	numRexExp := regexp.MustCompile(`^\d+$`)
	if !numRexExp.MatchString(length) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}
	len, err := strconv.Atoi(length)
	if err != nil {
		log.Fatal(err)
	}

	line := Parentheses{
		RandParentheses(len),
	}

	res, err := json.Marshal(line)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func percents(w http.ResponseWriter, r *http.Request) {
	m := CalculateBalancedPercentage()

	result := Percents{
		m[2],
		m[4],
		m[8],
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequests(adr string) {
	http.HandleFunc("/generate/", generate)
	log.Fatal(http.ListenAndServe(adr, nil))
}

func main() {
	handleRequests(":8080")
}
